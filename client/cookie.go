package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type cred struct {
	URL string `json:"url"`
	// Domain of the website. e.g. .tumblr.com
	Domain  string            `json:"domain"`
	Payload map[string]string `json:"payload"`
}

func toCookie(c *cred) *cookiejar.Jar {
	j, err := cookiejar.New(nil)
	if err != nil {
		return nil
	}

	u, err := url.Parse(c.URL)
	if err != nil {
		return nil
	}

	var cookies []*http.Cookie

	for n, v := range c.Payload {
		cookies = append(cookies, &http.Cookie{
			Domain:   c.Domain,
			Expires:  time.Now().Add(time.Hour),
			Secure:   true,
			HttpOnly: true,
			Path:     "/",
			Value:    v,
			Name:     n,
		})
	}

	j.SetCookies(u, cookies)
	return j
}

func loadFromFile(filePath string) *cred {
	d, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil
	}
	var res cred
	if err := json.Unmarshal(d, &res); err != nil {
		return nil
	}
	return &res
}
