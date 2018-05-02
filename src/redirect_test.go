package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func init() {
	whiteListedSuffix = "test.foo"
}

func TestHandle301(t *testing.T) {
	hostCheckEnabled = true

	redirectCode = 301
	redirectURL = "redirect.too/test"

	req, err := http.NewRequest("GET", "test.foo/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	execute(t, req, redirectCode, "https://"+redirectURL)
}

func TestHandle302(t *testing.T) {
	hostCheckEnabled = true
	whiteListedSuffix = "test.foo"
	redirectCode = 302
	redirectURL = "redirect.too/test"

	req, err := http.NewRequest("GET", "test.foo/test1", nil)
	req.RequestURI = "/test1"
	req.Host = "test.foo"
	if err != nil {
		t.Fatal(err)
	}

	execute(t, req, redirectCode, "https://"+"test.foo/test1")
}

func TestHandleBadHost(t *testing.T) {
	hostCheckEnabled = true
	redirectCode = 302
	redirectURL = "redirect.too/test"

	req, err := http.NewRequest("GET", "http://test.wrong/test", nil)
	req.URL = &url.URL{
		Scheme: "http",
		Host:   "test.wrong",
	}

	req.Header.Set("Host", "test.wrong")
	req.Host = "test.wrong"
	if err != nil {
		t.Fatal(err)
	}

	execute(t, req, redirectCode, "https://"+redirectURL)
}

func TestHandleBadHostHttp(t *testing.T) {
	hostCheckEnabled = true
	redirectCode = 302
	redirectURL = "redirect.too/test"
	redirectHTTP = true
	req, err := http.NewRequest("GET", "http://test.wrong/test", nil)
	req.URL = &url.URL{
		Scheme: "http",
		Host:   "test.wrong",
	}

	req.Header.Set("Host", "test.wrong")
	req.Host = "test.wrong"
	if err != nil {
		t.Fatal(err)
	}

	execute(t, req, redirectCode, "http://"+redirectURL)

	redirectHTTP = false
}

func TestHttp(t *testing.T) {
	hostCheckEnabled = true
	redirectCode = 302
	redirectURL = "redirect.too/test"
	redirectHTTP = true

	req, err := http.NewRequest("GET", "test.foo/test1", nil)
	req.RequestURI = "/test1"
	req.Host = "test.foo"
	if err != nil {
		t.Fatal(err)
	}

	execute(t, req, redirectCode, "http://"+"test.foo/test1")
	redirectHTTP = false

}

func execute(t *testing.T, req *http.Request, expectedCode int, expectedLocation string) {

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handle)

	handler.ServeHTTP(rr, req)

	assertEqual(t, rr.Code, expectedCode, "Redirect Code Incorrect")
	assertEqual(t, rr.Header().Get("Location"), expectedLocation, "Location was incorrect.")
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}
