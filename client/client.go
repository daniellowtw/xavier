package client

import (
	"net/http"
)

const (
	userAgent                 = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36"
	defaultCookieCredFilePath = ".cred.json"
)

type Option func(*http.Client)

func WithFileCookie(filePath string) Option {
	cred := loadFromFile(filePath)
	if cred == nil {
		panic("cannot load credentials from cookie")
	}
	j := toCookie(cred)
	if j == nil {
		panic("cannot load cookie jar")
	}
	return func(c *http.Client) {
		c.Jar = j
	}
}

func NewDefaultClient() *http.Client {
	return New(WithFileCookie(defaultCookieCredFilePath))
}

func New(opts ...Option) *http.Client {
	c := &http.Client{
		Transport: &myTransport{
			rt: http.DefaultTransport,
		},
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

type myTransport struct {
	rt http.RoundTripper
}

func (t *myTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("User-Agent", userAgent)
	return t.rt.RoundTrip(r)
}
