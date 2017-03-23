package client

import (
	"net/http"
)

const userAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:40.0) Gecko/20100101 Firefox/40.1"

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
