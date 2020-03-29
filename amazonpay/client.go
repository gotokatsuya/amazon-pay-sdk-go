package amazonpay

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/go-playground/form"
)

const (
	EndpointHostJP = "mws.amazonservices.jp"
)

const (
	EndpointPathReal    = "/OffAmazonPayments"
	EndpointPathSandbox = "/OffAmazonPayments_Sandbox"
)

const (
	Version = "2013-01-01"
)

// Client type
type Client struct {
	AccessKey        string
	SecretKey        string
	SellerID         string
	SignatureMethod  string
	SignatureVersion string
	Version          string
	EndpointHost     string
	EndpointPath     string

	endpoint   *url.URL
	httpClient *http.Client
}

// ClientOption type
type ClientOption func(*Client) error

// New returns a new pay client instance.
func New(accessKey, secretKey, sellerID string, options ...ClientOption) (*Client, error) {
	if accessKey == "" {
		return nil, errors.New("missing accessKey")
	}
	if secretKey == "" {
		return nil, errors.New("missing secretKey")
	}
	if sellerID == "" {
		return nil, errors.New("missing sellerID")
	}
	c := &Client{
		AccessKey:        accessKey,
		SecretKey:        secretKey,
		SellerID:         sellerID,
		SignatureMethod:  "HmacSHA256",
		SignatureVersion: "2",
		Version:          Version,
		EndpointHost:     EndpointHostJP,
		EndpointPath:     EndpointPathReal,
		httpClient:       http.DefaultClient,
	}
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	endpoint := "https://" + c.EndpointHost + c.EndpointPath + "/" + Version
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	c.endpoint = u
	return c, nil
}

// WithHTTPClient function
func WithHTTPClient(c *http.Client) ClientOption {
	return func(client *Client) error {
		client.httpClient = c
		return nil
	}
}

// WithEndpointHost function
func WithEndpointHost(endpointHost string) ClientOption {
	return func(client *Client) error {
		client.EndpointHost = endpointHost
		return nil
	}
}

// WithEndpointPath function
func WithEndpointPath(endpointPath string) ClientOption {
	return func(client *Client) error {
		client.EndpointPath = endpointPath
		return nil
	}
}

// WithSandbox function
func WithSandbox() ClientOption {
	return WithEndpointPath(EndpointPathSandbox)
}

// NewRequest method
func (c *Client) NewRequest(action string, body interface{}) (*http.Request, error) {
	values, err := form.NewEncoder().Encode(body)
	if err != nil {
		return nil, err
	}
	data := NewRequestValues(values)
	data.Set("AWSAccessKeyId", c.AccessKey)
	data.Set("Action", action)
	data.Set("SellerId", c.SellerID)
	data.Set("SignatureMethod", c.SignatureMethod)
	data.Set("SignatureVersion", c.SignatureVersion)
	data.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	data.Set("Version", c.Version)
	data.Set("Signature", c.makeSignature(http.MethodPost, data.RawEncode()))
	req, err := http.NewRequest(http.MethodPost, c.endpoint.String(), strings.NewReader(data.RawEncode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}

// https://m.media-amazon.com/images/G/09/AmazonPayments/Signature.pdf
func (c *Client) makeSignature(method string, requestValues string) string {
	key := []byte(c.SecretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(strings.Join([]string{method, c.EndpointHost, c.EndpointPath + "/" + c.Version, requestValues}, "\n")))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// Do method
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req.WithContext(ctx))
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
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			switch {
			case resp.StatusCode == http.StatusOK:
				if err := xml.Unmarshal(data, v); err != nil {
					return nil, err
				}
			default:
				var responseErr ResponseError
				if err := xml.Unmarshal(data, &responseErr); err == nil {
					return nil, responseErr
				}
			}
		}
	}
	return resp, nil
}

// -- request values encoder --

type requestValues struct {
	url.Values
}

func NewRequestValues(v url.Values) requestValues {
	return requestValues{v}
}

func (v requestValues) RawEncode() string {
	var buf strings.Builder
	keys := make([]string, 0, len(v.Values))
	for k := range v.Values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v.Values[k]
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(strings.Replace(url.QueryEscape(v), "+", "%20", -1))
		}
	}
	return buf.String()
}
