package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProxy(t *testing.T) {
	test := httptest.NewServer(http.HandlerFunc(proxyHandler))

	resp, err := http.Get(test.URL)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if _, err := io.ReadAll(resp.Body); err != nil {
		t.Log(err)
		t.Fail()
	}
	defer resp.Body.Close()
}
