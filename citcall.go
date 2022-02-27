package citcallgo

import (
	"net/http"
	"net/url"
)

const (
	defaultApiURL = "https://citcall.pub"
	apiVersion    = "v3"
)

type ApiKey string

type CitcallOption func(*Citcall)

type Citcall struct {
	httpClient *http.Client
	baseUrl    *url.URL

	apiUrl     string
	apiVersion string

	apiKey ApiKey
}

func New(apiKey ApiKey, opts ...CitcallOption) *Citcall {
	u, err := url.Parse(defaultApiURL)
	if err != nil {
		panic(err)
	}

	c := &Citcall{
		httpClient: http.DefaultClient,
		baseUrl:    u,
		apiUrl: defaultApiURL,
		apiVersion: apiVersion,
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