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

type CitcallOption func(*Citcall)

type Citcall struct {
	httpClient *http.Client
	baseUrl    *url.URL

	apiUrl     string
	apiVersion string
	citcallURL CitcallURL

	apiKey apiKey
}

func New(apiKey apiKey, opts ...CitcallOption) *Citcall {

	citcallURL := NewCitcallURL()

	u, err := url.Parse(citcallURL.DefaultApiURL)
	if err != nil {
		panic(err)
	}

	c := &Citcall{
		httpClient: http.DefaultClient,
		baseUrl:    u,
		apiUrl:     citcallURL.DefaultApiURL,
		apiVersion: citcallURL.ApiVersion,
		citcallURL: *citcallURL,
		apiKey:     apiKey,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// SetCustomApiURL overrides current Citcall instance API url
func (c *Citcall) SetCustomApiURL(url string) {
	c.apiUrl = url
}

// SetCustomApiVersion overrides current Citcall instance API version
func (c *Citcall) SetCustomApiVersion(version string) {
	c.apiVersion = version
}

// Initialize Citcall with custom API url
func WithCustomApiURL(url string) CitcallOption {
	return func(c *Citcall) {
		c.apiUrl = url
	}
}

// Initialize Citcall with custom API version
func WithCustomApiVersion(version string) CitcallOption {
	return func(c *Citcall) {
		c.apiVersion = version
	}
}

func (c *Citcall) request(ctx context.Context, method string, urlStr string, requestBody interface{}) (*http.Response, error) {
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
