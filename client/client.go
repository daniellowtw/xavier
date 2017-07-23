package client

import (
	"net/http"
)

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36"

func New() *http.Client {
	return &http.Client{
		Transport: &myTransport{
			rt: http.DefaultTransport,
		},
	}
}

type myTransport struct {
	rt http.RoundTripper
}

func (t *myTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("User-Agent", userAgent)
	return t.rt.RoundTrip(r)
}
