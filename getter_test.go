package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var siteToUrl map[string]string

func Test_getAndUpdateData_basicHtmlCase(t *testing.T) {
	siteToServer := setupServers()
	defer tearDownServers(siteToServer)
	itemData := makeBasicItemDataForHtmlParsing()
	getAndUpdateData(&itemData)

	AssertEqual(itemData.SiteHistory["htmlHandler1"].GetLatestPrice(), 150.00, t)
}

func Test_getAndUpdateData_basicJsonCase(t *testing.T) {
	siteToServer := setupServers()
	defer tearDownServers(siteToServer)
	itemData := makeBasicItemDataForJsonParsing()
	getAndUpdateData(&itemData)

	AssertEqual(itemData.SiteHistory["jsonHandler1"].GetLatestPrice(), 123.45, t)
}

func Test_getAndUpdateData_hasMultipleSources(t *testing.T) {
	siteToServer := setupServers()
	defer tearDownServers(siteToServer)
	itemData := makeBasicItemDataForMultipleParsing()
	getAndUpdateData(&itemData)

	AssertEqual(itemData.SiteHistory["htmlHandler2"].GetLatestPrice(), -1.0, t)
	AssertEqual(itemData.SiteHistory["htmlHandler3"].GetLatestPrice(), 105.00, t)
	AssertEqual(itemData.SiteHistory["jsonHandler2"].GetLatestPrice(), -1.0, t)
	AssertEqual(itemData.SiteHistory["jsonHandler3"].GetLatestPrice(), 45.00, t)
}

func Test_GetAndUpdateDataForItems_WithMultipleGoroutines(t *testing.T) {
	cases := []string{"with 1 goroutine", "with 2 goroutines", "with 3 goroutines", "with 4 goroutines", "with 5 goroutines"}

	for index, name := range cases {
		t.Run(name, func(t *testing.T) {
			siteToServer := setupServers()
			defer tearDownServers(siteToServer)

			items := []ItemData{makeBasicItemDataForHtmlParsing(), makeBasicItemDataForJsonParsing(), makeBasicItemDataForMultipleParsing()}
			itemsToUpdate := []*ItemData{}

			for i := range items {
				itemsToUpdate = append(itemsToUpdate, &items[i])
			}

			GetAndUpdateDataForItems(itemsToUpdate, index+1)
			AssertEqual(itemsToUpdate[0].SiteHistory["htmlHandler1"].GetLatestPrice(), 150.00, t)
			AssertEqual(itemsToUpdate[1].SiteHistory["jsonHandler1"].GetLatestPrice(), 123.45, t)
			AssertEqual(itemsToUpdate[2].SiteHistory["htmlHandler2"].GetLatestPrice(), -1.0, t)
			AssertEqual(itemsToUpdate[2].SiteHistory["htmlHandler3"].GetLatestPrice(), 105.00, t)
			AssertEqual(itemsToUpdate[2].SiteHistory["jsonHandler2"].GetLatestPrice(), -1.0, t)
			AssertEqual(itemsToUpdate[2].SiteHistory["jsonHandler3"].GetLatestPrice(), 45.00, t)
		})
	}
}

func setupServers() map[string]*httptest.Server {
	siteToUrl = make(map[string]string)
	siteToServer := make(map[string]*httptest.Server)

	s := httptest.NewServer(http.HandlerFunc(htmlHandler1))
	siteToUrl["htmlHandler1"] = s.URL
	siteToServer["htmlHandler1"] = s

	s = httptest.NewServer(http.HandlerFunc(jsonHandler1))
	siteToUrl["jsonHandler1"] = s.URL
	siteToServer["jsonHandler1"] = s

	s = httptest.NewServer(http.HandlerFunc(htmlHandler2))
	siteToUrl["htmlHandler2"] = s.URL
	siteToServer["htmlHandler2"] = s

	s = httptest.NewServer(http.HandlerFunc(jsonHandler2))
	siteToUrl["jsonHandler2"] = s.URL
	siteToServer["jsonHandler2"] = s

	s = httptest.NewServer(http.HandlerFunc(htmlHandler3))
	siteToUrl["htmlHandler3"] = s.URL
	siteToServer["htmlHandler3"] = s

	s = httptest.NewServer(http.HandlerFunc(jsonHandler3))
	siteToUrl["jsonHandler3"] = s.URL
	siteToServer["jsonHandler3"] = s

	return siteToServer
}

func tearDownServers(siteToServer map[string]*httptest.Server) {
	for _, server := range siteToServer {
		CloseServer(server)
	}
}

func htmlHandler1(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, `<html><head></head><body><span class="check">150.00</span></body></html>`)
}

func jsonHandler1(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, `{"item": "box", "price": 123.45}`)
}

func htmlHandler2(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, `<html><head></head><body><span class="check">101.00</span></body></html>`)
}

func jsonHandler2(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, `{"item": "box", "price": 13.00}`)
}

func htmlHandler3(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, `<html><head></head><body><span class="check">105.00</span></body></html>`)
}

func jsonHandler3(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, `{"item": "box", "price": 45.00}`)
}

func makeBasicItemDataForHtmlParsing() ItemData {
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

	return itemData
}

func makeBasicItemDataForJsonParsing() ItemData {
	itemSiteData := make([]ItemSiteData, 0)
	itemSiteData = append(itemSiteData, ItemSiteData{
		SiteId:     2,
		Name:       "jsonHandler1",
		ParserType: jsonParser,
		Path:       []interface{}{"price"},
		Url:        siteToUrl["jsonHandler1"]})

	itemData := ItemData{
		Name:        "Box",
		ItemId:      2,
		SiteData:    itemSiteData,
		SiteHistory: make(map[string]*SiteHistory)}

	itemData.SiteHistory["jsonHandler1"] = &SiteHistory{}

	return itemData
}

func makeBasicItemDataForMultipleParsing() ItemData {
	itemSiteData := make([]ItemSiteData, 0)

	itemSiteData = append(itemSiteData, ItemSiteData{
		SiteId:     3,
		Name:       "htmlHandler2",
		ParserType: htmlParser,
		Path:       []interface{}{".check1"},
		Url:        siteToUrl["htmlHandler2"]})

	itemSiteData = append(itemSiteData, ItemSiteData{
		SiteId:     4,
		Name:       "htmlHandler3",
		ParserType: htmlParser,
		Path:       []interface{}{".check"},
		Url:        siteToUrl["htmlHandler3"]})

	itemSiteData = append(itemSiteData, ItemSiteData{
		SiteId:     5,
		Name:       "jsonHandler2",
		ParserType: jsonParser,
		Path:       []interface{}{".check"},
		Url:        siteToUrl["jsonHandler2"]})

	itemSiteData = append(itemSiteData, ItemSiteData{
		SiteId:     6,
		Name:       "jsonHandler3",
		ParserType: jsonParser,
		Path:       []interface{}{"price"},
		Url:        siteToUrl["jsonHandler3"]})

	itemData := ItemData{
		Name:        "Box",
		ItemId:      3,
		SiteData:    itemSiteData,
		SiteHistory: make(map[string]*SiteHistory)}

	itemData.SiteHistory["htmlHandler2"] = &SiteHistory{}
	itemData.SiteHistory["htmlHandler3"] = &SiteHistory{}
	itemData.SiteHistory["jsonHandler2"] = &SiteHistory{}
	itemData.SiteHistory["jsonHandler3"] = &SiteHistory{}

	return itemData
}
