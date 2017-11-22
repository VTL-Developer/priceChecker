package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCostHtml_EmptyCase(t *testing.T) {
	testGetCostHtml(t, []byte(`<body>123.45</body>`), []interface{}{}, 123.45)
}

func TestGetCostHtml_ParseHtml(t *testing.T) {
	testGetCostHtml(t, []byte(`<body><div><span class="cost">123.45</span></div></body>`), []interface{}{"div > span"}, 123.45)
}

func TestGetCostHtml_ParseClass(t *testing.T) {
	testGetCostHtml(t, []byte(`<body><div><span class="cost">123.45</span></div></body>`), []interface{}{".cost"}, 123.45)
}

func TestGetCostHtml_ParseHtmlInParts(t *testing.T) {
	testGetCostHtml(t, []byte(`<body><div><span class="cost">123.45</span></div></body>`), []interface{}{"div", "span"}, 123.45)
}

func Test_getDataAsHtml(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(httpHandler))
	defer CloseServer(s)
	rsp, _ := http.Get(s.URL)

	parser := HtmlParser{}
	document, err := parser.GetData(rsp)

	AssertNotError(err, t)

	if convertedDocument, ok := (*document).(goquery.Document); !ok {
		t.Errorf("Conversion failed.\nOriginal: %q\nConverted: %q", document, convertedDocument)
	} else {
		span := convertedDocument.Find("span")
		AssertEqual(span.Text(), "hello", t)
		AssertTrue(span.HasClass("check"), t)
	}
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

func testGetCostHtml(t *testing.T, htmlBody []byte, path []interface{}, expected float64) {
	r := bytes.NewReader(htmlBody)
	document, err := goquery.NewDocumentFromReader(r)

	AssertNotError(err, t)

	var interfaceObj interface{} = *document

	parser := HtmlParser{}
	got := parser.GetCost(&interfaceObj, path)
	AssertEqual(got, expected, t)
}
