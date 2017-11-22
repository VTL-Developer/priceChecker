package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getHttpBodyResponse_NilHeaders(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(httpHandler))
	defer CloseServer(s)
	rsp, err := getHttpBodyResponse(s.URL, nil)

	if rsp != nil {
		defer rsp.Body.Close()
	}

	AssertNotError(err, t)
	AssertEqual(rsp.Request.URL.String(), s.URL, t)
	AssertEqual(len(rsp.Request.Header), 0, t)
}

func Test_getHttpBodyResponse_EmptyHeaders(t *testing.T) {
	headers := make(map[string]string)
	s := httptest.NewServer(http.HandlerFunc(httpHandler))
	defer CloseServer(s)
	rsp, err := getHttpBodyResponse(s.URL, headers)

	if rsp != nil {
		defer rsp.Body.Close()
	}

	AssertNotError(err, t)

	AssertEqual(rsp.Request.URL.String(), s.URL, t)
	AssertEqual(len(rsp.Request.Header), 0, t)
}

func Test_getHttpBodyResponse_WithHeaders(t *testing.T) {
	headers := make(map[string]string)
	headers["hello"] = "world"
	headers["another"] = "item"

	s := httptest.NewServer(http.HandlerFunc(httpHandler))
	defer CloseServer(s)
	rsp, err := getHttpBodyResponse(s.URL, headers)

	if rsp != nil {
		defer rsp.Body.Close()
	}

	AssertNotError(err, t)

	AssertEqual(rsp.Request.URL.String(), s.URL, t)
	AssertTrue(len(rsp.Request.Header) >= 2, t)
	AssertEqual(rsp.Request.Header.Get("hello"), "world", t)
	AssertEqual(rsp.Request.Header.Get("another"), "item", t)
}
