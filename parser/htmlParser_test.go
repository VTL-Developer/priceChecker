package main

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
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

func testGetCostHtml(t *testing.T, htmlBody []byte, path []interface{}, expected float64) {
	r := bytes.NewReader(htmlBody)
	document, err := goquery.NewDocumentFromReader(r)

	if err != nil {
		t.Errorf("Error thrown: %v", err)
		return
	}

	var interfaceObj interface{} = *document

	got := GetCostHtml(&interfaceObj, path)
	if got != expected {
		t.Errorf("GetCost returned %q, should be %q", got, expected)
	}
}
