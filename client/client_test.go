package client

import (
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestHeader(t *testing.T) {
	addrC := make(chan string)
	go createBackend(addrC)
	port := <-addrC
	client := New()
	resp, err := client.Get(fmt.Sprintf("http://localhost:%s", port))
	if err != nil {
		log.Fatalf(err.Error())
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	//println(string(bytes))
	assert.Contains(t, string(bytes), userAgent, "User agent is wrong")
}

func createBackend(addr chan string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(spew.Sdump(r.Header)))
	}
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf(err.Error())
	}
	_, port, _ := net.SplitHostPort(lis.Addr().String())
	addr <- port
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf("localhost:%s", port), nil)
}
