package api

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

// uses google service to get favicon

const (
	googleFavIconService = `https://www.google.com/s2/favicons?domain=`
)

func getFavIcon(url string) (string, error) {
	if url == "" {
		return "", fmt.Errorf("empty url")
	}
	resp, err := http.Get(googleFavIconService + url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	dd, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if len(dd) == 0 {
		return "", fmt.Errorf("fav: no data")
	}
	return fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(dd)), nil
}
