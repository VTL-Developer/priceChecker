package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getDataAsJson_With_PriceHistoryJson(t *testing.T) {
	var rsp http.Response
	rsp.Body = &MockClosingBuffer{bytes.NewBufferString(`{"price":"54.05", "datetime":"2017-10-24T22:15:21.000000Z"}`)}
	jsonBody, err := getDataAsJson(&rsp)

	AssertNotError(err, t)

	convertedJsonBody, ok := (*jsonBody).(map[string]interface{})

	if !ok {
		t.Errorf("Conversion was not successful.\nOriginal: %q\nConverted: %q", jsonBody, convertedJsonBody)
		return
	}

	AssertEqual(convertedJsonBody["price"], "54.05", t)
	AssertEqual(convertedJsonBody["datetime"], "2017-10-24T22:15:21.000000Z", t)
}

func Test_getDataAsJson_With_BadJSON(t *testing.T) {
	var rsp http.Response
	rsp.Body = &MockClosingBuffer{bytes.NewBufferString(`{"price":"54.05", "datetime":"2017-10-24T22:15:21.000000Z"`)}
	_, err := getDataAsJson(&rsp)

	AssertError(err, t)
}

func Test_getDataAsHtml(t *testing.T) {
	var convertedDocument goquery.Document
	s := httptest.NewServer(http.HandlerFunc(httpHandler))
	rsp, _ := http.Get(s.URL)

	document, err := getDataAsHtml(rsp)

	AssertNotError(err, t)

	convertedDocument, ok := (*document).(goquery.Document)

	if !ok {
		t.Errorf("Conversion failed.\nOriginal: %q\nConverted: %q", document, convertedDocument)
		return
	}

	span := convertedDocument.Find("span")

	AssertEqual(span.Text(), "hello", t)
	AssertTrue(span.HasClass("check"), t)
}

func Test_getHttpBodyResponse_NilHeaders(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(httpHandler))
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

func httpHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, `<html><head></head><body><span class="check">hello</span></body></html>`)
}

type MockClosingBuffer struct {
	*bytes.Buffer
}

func (mcb *MockClosingBuffer) Close() error {
	return nil
}
