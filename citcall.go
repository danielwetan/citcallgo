package citcallgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type apiKey string

func (k apiKey) String() string {
	return string(k)
}

type citcallOption func(*citcall)

type citcall struct {
	httpClient *http.Client
	baseUrl    *url.URL

	apiUrl     string
	apiVersion string
	citcallURL citcallURL

	apiKey apiKey
}

// Initialize new citcallgo instance
func New(apiKey apiKey, opts ...citcallOption) *citcall {

	citcallURL := newCitcallURL()

	u, err := url.Parse(citcallURL.defaultApiURL)
	if err != nil {
		panic(err)
	}

	c := &citcall{
		httpClient: http.DefaultClient,
		baseUrl:    u,
		apiUrl:     citcallURL.defaultApiURL,
		apiVersion: citcallURL.apiVersion,
		citcallURL: *citcallURL,
		apiKey:     apiKey,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Initialize citcallgo with custom API url
func WithCustomApiURL(url string) citcallOption {
	return func(c *citcall) {
		c.apiUrl = url
	}
}

// Initialize citcallgo with custom API version
func WithCustomApiVersion(version string) citcallOption {
	return func(c *citcall) {
		c.apiVersion = version
	}
}

func (c *citcall) request(ctx context.Context, method string, urlStr string, requestBody interface{}) (*http.Response, error) {
	u, err := c.baseUrl.Parse(fmt.Sprintf("%s/%s", c.apiVersion, urlStr))
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if requestBody != nil {
		body, err := json.Marshal(requestBody)
		if err != nil {
			return nil, err
		}

		buf = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Apikey %s", c.apiKey.String()))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		var apiErr Error
		err = json.NewDecoder(res.Body).Decode(&apiErr)
		if err != nil {
			return nil, err
		}

		return nil, &apiErr
	}

	return res, nil
}
