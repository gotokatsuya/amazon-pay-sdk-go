package signing

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

const (
	Algorithm = "AMZN-PAY-RSASSA-PSS"
)

// CanonicalRequest =
//  HTTPRequestMethod + '\n' +
//  CanonicalURI + '\n' +
//  CanonicalQueryString + '\n' +
//  CanonicalHeaders + '\n' +
//  SignedHeaders + '\n' +
//  HexEncode(Hash(RequestPayload))
func CanonicalRequest(r *http.Request) (string, error) {
	data, err := RequestPayload(r)
	if err != nil {
		return "", err
	}
	hexencode, err := HexEncodeSHA256Hash(data)
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", r.Method, CanonicalURI(r), CanonicalQueryString(r), CanonicalHeaders(r), SignedHeaders(r), hexencode), err
}

// CanonicalURI returns request uri
func CanonicalURI(r *http.Request) string {
	pattens := strings.Split(r.URL.Path, "/")
	var uri []string
	for _, v := range pattens {
		switch v {
		case "":
			continue
		case ".":
			continue
		case "..":
			if len(uri) > 0 {
				uri = uri[:len(uri)-1]
			}
		default:
			uri = append(uri, url.QueryEscape(v))
		}
	}
	urlpath := "/" + strings.Join(uri, "/")
	r.URL.Path = strings.Replace(urlpath, "+", "%20", -1)
	return fmt.Sprintf("%s", r.URL.Path)
}

func CanonicalQueryString(r *http.Request) string {
	var a []string
	for key, value := range r.URL.Query() {
		k := url.QueryEscape(key)
		for _, v := range value {
			var kv string
			if v == "" {
				kv = k
			} else {
				kv = fmt.Sprintf("%s=%s", k, url.QueryEscape(v))
			}
			a = append(a, strings.Replace(kv, "+", "%20", -1))
		}
	}
	sort.Strings(a)
	return fmt.Sprintf("%s", strings.Join(a, "&"))
}

func CanonicalHeaders(r *http.Request) string {
	var a []string
	for key, value := range r.Header {
		sort.Strings(value)
		var q []string
		for _, v := range value {
			q = append(q, trimString(v))
		}
		a = append(a, strings.ToLower(key)+":"+strings.Join(q, ","))
	}
	sort.Strings(a)
	return fmt.Sprintf("%s\n", strings.Join(a, "\n"))
}

func SignedHeaders(r *http.Request) string {
	var a []string
	for key := range r.Header {
		a = append(a, strings.ToLower(key))
	}
	sort.Strings(a)
	return fmt.Sprintf("%s", strings.Join(a, ";"))
}

func RequestPayload(r *http.Request) ([]byte, error) {
	if r.Body == nil {
		return []byte(""), nil
	}
	b, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	return b, err
}

func StringToSign(canonicalRequest string) (string, error) {
	hexencode, err := HexEncodeSHA256Hash([]byte(canonicalRequest))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s\n%s", Algorithm, hexencode), nil
}

func Sign(privateKeyData []byte, stringToSign string) (string, error) {
	block, _ := pem.Decode(privateKeyData)
	if block == nil {
		return "", errors.New("invalid private key data")
	}
	keyIF, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	key, ok := keyIF.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("not RSA private key")
	}
	hashed := sha256.Sum256([]byte(stringToSign))
	signature, err := rsa.SignPSS(rand.Reader, key, crypto.SHA256, hashed[:], &rsa.PSSOptions{
		SaltLength: 20,
	})
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

func HexEncodeSHA256Hash(body []byte) (string, error) {
	hash := sha256.New()
	if body == nil {
		body = []byte("")
	}
	_, err := hash.Write(body)
	return fmt.Sprintf("%x", hash.Sum(nil)), err
}

func AuthHeaderValue(publicKeyID, signedHeaders, signature string) string {
	return fmt.Sprintf("%s PublicKeyId=%s, SignedHeaders=%s, Signature=%s", Algorithm, publicKeyID, signedHeaders, signature)
}

func trimString(s string) string {
	var trimedString []byte
	inQuote := false
	var lastChar byte
	s = strings.TrimSpace(s)
	for _, v := range []byte(s) {
		if byte(v) == byte('"') {
			inQuote = !inQuote
		}
		if lastChar == byte(' ') && byte(v) == byte(' ') && !inQuote {
			continue
		}
		trimedString = append(trimedString, v)
		lastChar = v
	}
	return string(trimedString)
}
