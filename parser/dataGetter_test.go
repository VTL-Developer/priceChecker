package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"
)

func Test_getDataAsJson_With_PriceHistoryJson(t *testing.T) {
	var rsp http.Response
	rsp.Body = &MockClosingBuffer{bytes.NewBufferString(`{"price":"54.05", "datetime":"2017-10-24T22:15:21.000000Z"}`)}
	jsonBody, err := getDataAsJson(&rsp)

	if err != nil {
		t.Errorf("Error occured: %v", err)
		return
	}

	convertedJsonBody, ok := (*jsonBody).(map[string]interface{})

	if !ok {
		t.Errorf("Conversion was not successful.\nOriginal: %q\nConverted: %q", jsonBody, convertedJsonBody)
		return
	}

	if convertedJsonBody["price"] != "54.05" {
		t.Errorf("JSON body price is %q, should be %q", convertedJsonBody["price"], "54.05")
	}

	if convertedJsonBody["datetime"] != "2017-10-24T22:15:21.000000Z" {
		t.Errorf("JSON body datetime is %q, should be %q", convertedJsonBody["datetime"], "2017-10-24T22:15:21.000000Z")
	}
}

func Test_getDataAsJson_With_BadJSON(t *testing.T) {
	var rsp http.Response
	rsp.Body = &MockClosingBuffer{bytes.NewBufferString(`{"price":"54.05", "datetime":"2017-10-24T22:15:21.000000Z"`)}
	_, err := getDataAsJson(&rsp)

	if err == nil {
		t.Errorf("Error should have been thrown")
	}
}

func Test_getDataAsHtml(t *testing.T) {
	var convertedDocument goquery.Document
	s := httptest.NewServer(http.HandlerFunc(httpHandler))
	rsp, _ := http.Get(s.URL)

	document, err := getDataAsHtml(rsp)

	if err != nil {
		t.Errorf("Error occured: %v", err)
		return
	}

	convertedDocument, ok := (*document).(goquery.Document)

	if !ok {
		t.Errorf("Conversion failed.\nOriginal: %q\nConverted: %q", document, convertedDocument)
		return
	}

	span := convertedDocument.Find("span")

	if span.Text() != "hello" {
		t.Errorf("Span text found is %q, should be %q", span.Text(), "hello")
	}

	if !span.HasClass("check") {
		t.Error("Span does not have class \"check\"")
	}
}

func Test_getHttpBodyResponse_NilHeaders(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(httpHandler))
	rsp, err := getHttpBodyResponse(s.URL, nil)

	if rsp != nil {
		defer rsp.Body.Close()
	}

	if err != nil {
		t.Errorf("Error occured: %v", err)
		return
	}

	if rsp.Request.URL.String() != s.URL {
		t.Errorf("Request URL was %q, should be %q", rsp.Request.URL, s.URL)
	}

	if len(rsp.Request.Header) != 0 {
		t.Errorf("Request heaer should be empty, but has %q", rsp.Request.Header)
	}
}

func Test_getHttpBodyResponse_EmptyHeaders(t *testing.T) {
	headers := make(map[string]string)
	s := httptest.NewServer(http.HandlerFunc(httpHandler))
	rsp, err := getHttpBodyResponse(s.URL, headers)

	if rsp != nil {
		defer rsp.Body.Close()
	}

	if err != nil {
		t.Errorf("Error occured: %v", err)
		return
	}

	if rsp.Request.URL.String() != s.URL {
		t.Errorf("Request URL was %q, should be %q", rsp.Request.URL, s.URL)
	}

	if len(rsp.Request.Header) != 0 {
		t.Errorf("Request heaer should be empty, but has %q", rsp.Request.Header)
	}
}

func Test_getHttpBodyResponse_WithHeaders(t *testing.T) {
	headers := make(map[string]string)
	headers["hello"] = "world"
	headers["another"] = "item"

	s := httptest.NewServer(http.HandlerFunc(httpHandler))
	rsp, err := getHttpBodyResponse(s.URL, headers)

	if rsp != nil {
		defer rsp.Body.Close()
	}

	if err != nil {
		t.Errorf("Error occured: %v", err)
		return
	}

	if rsp.Request.URL.String() != s.URL {
		t.Errorf("Request URL was %q, should be %q", rsp.Request.URL, s.URL)
	}

	if len(rsp.Request.Header) < 2 {
		t.Errorf("Request heaer should have two items, but has %q item(s) with: %q", len(rsp.Request.Header), rsp.Request.Header)
	}

	if rsp.Request.Header.Get("hello") != "world" {
		t.Errorf("Request headers should have 'hello: world' but has 'hello: %v'", rsp.Request.Header.Get("hello"))
	}

	if rsp.Request.Header.Get("another") != "item" {
		t.Errorf("Request headers should have 'another: item' but has 'another: %v'", rsp.Request.Header.Get("another"))
	}
}

func httpHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, `<html><head></head><body><span class="check">hello</span></body></html>`, path.Base(r.URL.Path))
}

type MockClosingBuffer struct {
	*bytes.Buffer
}

func (mcb *MockClosingBuffer) Close() error {
	return nil
}
