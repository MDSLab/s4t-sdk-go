package s4t


import (
	"net/http"
	"time"
)

type Client struct {
	HTTPClient *http.Client
	AuthToken string
	Endpoint string
	Timeout time.Duration
}


type ClientOption func( *Client )


func NewClient(endpoint string, options ...ClientOption) *Client {
	c := &Client{
		HTTPClient: &http.Client{},
		Endpoint: endpoint,
		Timeout: time.Second * 30,
	}

	for _, option := range options {
		option(c)
	}

	c.HTTPClient.Timeout = c.Timeout

	return c
}


func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.Timeout = timeout
	}
}
