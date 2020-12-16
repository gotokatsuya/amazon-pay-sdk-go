package amazonpay

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"time"

	"github.com/rs/xid"

	"github.com/gotokatsuya/amazon-pay-sdk-go/amazonpay/signing"
)

const (
	SDKVersion = "2.2.1"
)

const (
	APIVersion = "v2"
)

var (
	RegionMap = map[string]string{
		"eu": "eu",
		"de": "eu",
		"uk": "eu",
		"us": "na",
		"na": "na",
		"jp": "jp",
	}
	RegionHostMap = map[string]string{
		"eu": "pay-api.amazon.eu",
		"na": "pay-api.amazon.com",
		"jp": "pay-api.amazon.jp",
	}
)

// Client type
type Client struct {
	PublicKeyID string
	PrivateKey  []byte
	Region      string
	Sandbox     bool
	HTTPClient  *http.Client

	endpoint *url.URL
}

// New returns a new pay client instance.
func New(publicKeyID string, privateKey []byte, region string, sandbox bool, httpClient *http.Client) (*Client, error) {
	if publicKeyID == "" {
		return nil, errors.New("missing publicKeyID")
	}
	if privateKey == nil {
		return nil, errors.New("missing  privateKey")
	}
	if region == "" {
		return nil, errors.New("missing region")
	}
	c := &Client{
		PublicKeyID: publicKeyID,
		PrivateKey:  privateKey,
		Region:      region,
		Sandbox:     sandbox,
		HTTPClient:  httpClient,
	}
	endpointURL := c.createEndpointURL()
	u, err := url.Parse(endpointURL)
	if err != nil {
		return nil, err
	}
	c.endpoint = u
	return c, nil
}

func (c *Client) createEndpointURL() string {
	modePath := "live"
	if c.Sandbox {
		modePath = "sandbox"
	}
	host := RegionHostMap[RegionMap[c.Region]]
	return "https://" + host + "/" + modePath + "/"
}

// NewRequest method
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	u, err := c.endpoint.Parse(path)
	if err != nil {
		return nil, err
	}

	var reqBody io.ReadWriter
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, u.String(), reqBody)
	if err != nil {
		return nil, err
	}

	if method == http.MethodPost {
		req.Header.Set("x-amz-pay-idempotency-key", xid.New().String())
	}
	req.Header.Set("x-amz-pay-region", c.Region)
	req.Header.Set("x-amz-pay-host", RegionHostMap[RegionMap[c.Region]])
	req.Header.Set("x-amz-pay-date", time.Now().UTC().Format("20060102T150405Z"))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("amazon-pay-api-sdk-go/%s (GO/%s)", SDKVersion, runtime.Version()))

	canonicalRequest, err := signing.CanonicalRequest(req)
	if err != nil {
		return nil, err
	}
	stringToSign, err := signing.StringToSign(canonicalRequest)
	if err != nil {
		return nil, err
	}
	signature, err := signing.Sign(c.PrivateKey, stringToSign)
	if err != nil {
		return nil, err
	}
	signedHeaders := signing.SignedHeaders(req)
	authValue := signing.AuthHeaderValue(c.PublicKeyID, signedHeaders, signature)
	req.Header.Set("Authorization", authValue)

	return req, nil
}

// Do method
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req.WithContext(ctx))
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	defer resp.Body.Close()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
				return resp, err
			}
		}
	}
	return resp, nil
}
