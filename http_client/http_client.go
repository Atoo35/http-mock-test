// package http_client

// import (
// 	"net/http"
// 	"time"
// )

// var Client HTTPClient = &HttpClient{}

// type HttpClient struct {
// 	Client        *http.Client
// 	Transport     http.RoundTripper
// 	CheckRedirect func(req *http.Request, via []*http.Request) error
// 	Jar           http.CookieJar
// 	Timeout       time.Duration
// }

// type HTTPClient interface {
// 	Do(req *http.Request) (*http.Response, error)
// 	New(opts ...Option)
// 	WithTransport(transport *http.Transport) Option
// }

// type Option func(*HttpClient)

// func (h *HttpClient) Do(req *http.Request) (*http.Response, error) {
// 	if h.Client != nil {
// 		return h.Client.Do(req)
// 	}
// 	return nil, nil // Handle the case when the embedded Client is not set
// }

// func (h *HttpClient) New(opts ...Option) {
// 	s := &HttpClient{}
// 	for _, opt := range opts {
// 		opt(s)
// 	}
// 	if s.Client == nil {
// 		s.Client = &http.Client{}
// 	}
// }

// func (h *HttpClient) WithTransport(transport *http.Transport) Option {
// 	return func(s *HttpClient) {
// 		s.Transport = transport
// 		s.Client = &http.Client{Transport: transport}
// 	}
// }

// // package http_client

// // import (
// // 	"net/http"
// // 	"time"
// // )

// // var Client HTTPClient = &HttpClient{}

// // type HttpClient struct {
// // 	Transport http.RoundTripper
// // 	CheckRedirect func(req *http.Request, via []*http.Request) error
// // 	Jar http.CookieJar
// // 	Timeout time.Duration
// // }

// // type HTTPClient interface {
// // 	Do(req *http.Request) (*http.Response, error)
// // 	New(opts ...Option) *http.Client
// // 	WithTransport(transport *http.Transport) Option
// // }

// // type Option func(*http.Client)

// // func (h *HttpClient) New(opts ...Option) *http.Client {
// // 	s := &http.Client{}
// // 	for _, opt := range opts {
// // 		opt(s)
// // 	}
// // 	return s
// // }

// // func (h *HttpClient) WithTransport(transport *http.Transport) Option {
// // 	return func(s *http.Client) {
// // 		s.Transport = transport
// // 	}
// // }

package http_client

import (
	"net/http"
	"time"
)

var Client HTTPClient = newRealHTTPClient()

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RealHTTPClient struct {
	Client    *http.Client
	Transport http.RoundTripper
	Timeout   time.Duration
}

func (c *RealHTTPClient) Do(req *http.Request) (*http.Response, error) {
	c.Client.Transport = c.Transport
	c.Client.Timeout = c.Timeout
	return c.Client.Do(req)
}

func newRealHTTPClient() *RealHTTPClient {
	return &RealHTTPClient{
		Client: &http.Client{},
	}
}

func (c *RealHTTPClient) WithTransport(transport http.RoundTripper) *RealHTTPClient {
	c.Transport = transport
	return c
}

func (c *RealHTTPClient) WithTimeout(timeout time.Duration) *RealHTTPClient {
	c.Timeout = timeout
	return c
}
