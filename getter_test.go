package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var siteToUrl map[string]string

func Test_getAndUpdateData_basicHtmlCase(t *testing.T) {
	setup_server()
	itemSiteData := make([]ItemSiteData, 0)
	itemSiteData = append(itemSiteData, ItemSiteData{
		SiteId:     1,
		Name:       "htmlHandler1",
		ParserType: htmlParser,
		Path:       []interface{}{".check"},
		Url:        siteToUrl["htmlHandler1"]})

	itemData := ItemData{
		Name:        "Box",
		ItemId:      1,
		SiteData:    itemSiteData,
		SiteHistory: make(map[string]*SiteHistory)}

	itemData.SiteHistory["htmlHandler1"] = &SiteHistory{}
	getAndUpdateData(&itemData)

	AssertEqual(itemData.SiteHistory["htmlHandler1"].GetLatestPrice(), 150.00, t)
}

func Test_getAndUpdateData_basicJsonCase(t *testing.T) {
	setup_server()
	itemSiteData := make([]ItemSiteData, 0)
	itemSiteData = append(itemSiteData, ItemSiteData{
		SiteId:     1,
		Name:       "jsonHandler1",
		ParserType: jsonParser,
		Path:       []interface{}{"price"},
		Url:        siteToUrl["jsonHandler1"]})

	itemData := ItemData{
		Name:        "Box",
		ItemId:      1,
		SiteData:    itemSiteData,
		SiteHistory: make(map[string]*SiteHistory)}

	itemData.SiteHistory["jsonHandler1"] = &SiteHistory{}
	getAndUpdateData(&itemData)

	AssertEqual(itemData.SiteHistory["jsonHandler1"].GetLatestPrice(), 123.00, t)
}

func setup_server() {
	siteToUrl = make(map[string]string)
	s := httptest.NewServer(http.HandlerFunc(htmlHandler1))
	siteToUrl["htmlHandler1"] = s.URL

	s = httptest.NewServer(http.HandlerFunc(jsonHandler1))
	siteToUrl["jsonHandler1"] = s.URL
}

func htmlHandler1(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, `<html><head></head><body><span class="check">150.00</span></body></html>`)
}

func jsonHandler1(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, `{"item": "box", "price": 123.00}`)
}
