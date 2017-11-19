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
	itemData := makeBasicItemDataForHtmlParsing()
	getAndUpdateData(&itemData)

	AssertEqual(itemData.SiteHistory["htmlHandler1"].GetLatestPrice(), 150.00, t)
}

func Test_getAndUpdateData_basicJsonCase(t *testing.T) {
	setup_server()
	itemData := makeBasicItemDataForJsonParsing()
	getAndUpdateData(&itemData)

	AssertEqual(itemData.SiteHistory["jsonHandler1"].GetLatestPrice(), 123.45, t)
}

func Test_getAndUpdateData_hasMultipleSources(t *testing.T) {
	setup_server()
	itemData := makeBasicItemDataForMultipleParsing()
	getAndUpdateData(&itemData)

	AssertEqual(itemData.SiteHistory["htmlHandler1"].GetLatestPrice(), 150.00, t)
	AssertEqual(itemData.SiteHistory["htmlHandler2"].GetLatestPrice(), -1.0, t)
	AssertEqual(itemData.SiteHistory["jsonHandler1"].GetLatestPrice(), 123.45, t)
	AssertEqual(itemData.SiteHistory["jsonHandler2"].GetLatestPrice(), -1.0, t)
}

func Test_GetAndUpdateDataForItems_With1Goroutines(t *testing.T) {
	setup_server()

	items := []ItemData{makeBasicItemDataForHtmlParsing(), makeBasicItemDataForJsonParsing(), makeBasicItemDataForMultipleParsing()}
	itemsToUpdate := []*ItemData{}

	for _, item := range items {
		itemsToUpdate = append(itemsToUpdate, &item)
	}

	GetAndUpdateDataForItems(itemsToUpdate, 1)

	AssertEqual(itemsToUpdate[0].SiteHistory["htmlHandler1"].GetLatestPrice(), 150.00, t)
	AssertEqual(itemsToUpdate[1].SiteHistory["jsonHandler1"].GetLatestPrice(), 123.45, t)
	AssertEqual(itemsToUpdate[2].SiteHistory["htmlHandler1"].GetLatestPrice(), 150.00, t)
	AssertEqual(itemsToUpdate[2].SiteHistory["htmlHandler2"].GetLatestPrice(), -1.0, t)
	AssertEqual(itemsToUpdate[2].SiteHistory["jsonHandler1"].GetLatestPrice(), 123.45, t)
	AssertEqual(itemsToUpdate[2].SiteHistory["jsonHandler2"].GetLatestPrice(), -1.0, t)
}

func Test_GetAndUpdateDataForItems_With2Goroutines(t *testing.T) {
	setup_server()

	items := []ItemData{makeBasicItemDataForHtmlParsing(), makeBasicItemDataForJsonParsing(), makeBasicItemDataForMultipleParsing()}
	itemsToUpdate := []*ItemData{}

	for _, item := range items {
		itemsToUpdate = append(itemsToUpdate, &item)
	}

	GetAndUpdateDataForItems(itemsToUpdate, 2)

	AssertEqual(itemsToUpdate[0].SiteHistory["htmlHandler1"].GetLatestPrice(), 150.00, t)
	AssertEqual(itemsToUpdate[1].SiteHistory["jsonHandler1"].GetLatestPrice(), 123.45, t)
	AssertEqual(itemsToUpdate[2].SiteHistory["htmlHandler1"].GetLatestPrice(), 150.00, t)
	AssertEqual(itemsToUpdate[2].SiteHistory["htmlHandler2"].GetLatestPrice(), -1.0, t)
	AssertEqual(itemsToUpdate[2].SiteHistory["jsonHandler1"].GetLatestPrice(), 123.45, t)
	AssertEqual(itemsToUpdate[2].SiteHistory["jsonHandler2"].GetLatestPrice(), -1.0, t)
}

func Test_GetAndUpdateDataForItems_With3Goroutines(t *testing.T) {
	setup_server()

	items := []ItemData{makeBasicItemDataForHtmlParsing(), makeBasicItemDataForJsonParsing(), makeBasicItemDataForMultipleParsing()}
	itemsToUpdate := []*ItemData{}

	for _, item := range items {
		itemsToUpdate = append(itemsToUpdate, &item)
	}

	GetAndUpdateDataForItems(itemsToUpdate, 3)

	AssertEqual(itemsToUpdate[0].SiteHistory["htmlHandler1"].GetLatestPrice(), 150.00, t)
	AssertEqual(itemsToUpdate[1].SiteHistory["jsonHandler1"].GetLatestPrice(), 123.45, t)
	AssertEqual(itemsToUpdate[2].SiteHistory["htmlHandler1"].GetLatestPrice(), 150.00, t)
	AssertEqual(itemsToUpdate[2].SiteHistory["htmlHandler2"].GetLatestPrice(), -1.0, t)
	AssertEqual(itemsToUpdate[2].SiteHistory["jsonHandler1"].GetLatestPrice(), 123.45, t)
	AssertEqual(itemsToUpdate[2].SiteHistory["jsonHandler2"].GetLatestPrice(), -1.0, t)
}

func Test_GetAndUpdateDataForItems_With4Goroutines(t *testing.T) {
	setup_server()

	items := []ItemData{makeBasicItemDataForHtmlParsing(), makeBasicItemDataForJsonParsing(), makeBasicItemDataForMultipleParsing()}
	itemsToUpdate := []*ItemData{}

	for _, item := range items {
		itemsToUpdate = append(itemsToUpdate, &item)
	}

	GetAndUpdateDataForItems(itemsToUpdate, 4)

	AssertEqual(itemsToUpdate[0].SiteHistory["htmlHandler1"].GetLatestPrice(), 150.00, t)
	AssertEqual(itemsToUpdate[1].SiteHistory["jsonHandler1"].GetLatestPrice(), 123.45, t)
	AssertEqual(itemsToUpdate[2].SiteHistory["htmlHandler1"].GetLatestPrice(), 150.00, t)
	AssertEqual(itemsToUpdate[2].SiteHistory["htmlHandler2"].GetLatestPrice(), -1.0, t)
	AssertEqual(itemsToUpdate[2].SiteHistory["jsonHandler1"].GetLatestPrice(), 123.45, t)
	AssertEqual(itemsToUpdate[2].SiteHistory["jsonHandler2"].GetLatestPrice(), -1.0, t)
}

func setup_server() {
	siteToUrl = make(map[string]string)
	s := httptest.NewServer(http.HandlerFunc(htmlHandler1))
	siteToUrl["htmlHandler1"] = s.URL

	s = httptest.NewServer(http.HandlerFunc(jsonHandler1))
	siteToUrl["jsonHandler1"] = s.URL

	s = httptest.NewServer(http.HandlerFunc(htmlHandler2))
	siteToUrl["htmlHandler2"] = s.URL

	s = httptest.NewServer(http.HandlerFunc(jsonHandler2))
	siteToUrl["jsonHandler2"] = s.URL

	s = httptest.NewServer(http.HandlerFunc(htmlHandler3))
	siteToUrl["htmlHandler3"] = s.URL

	s = httptest.NewServer(http.HandlerFunc(jsonHandler3))
	siteToUrl["jsonHandler3"] = s.URL
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

	return itemData
}

func makeBasicItemDataForMultipleParsing() ItemData {
	itemSiteData := make([]ItemSiteData, 0)
	itemSiteData = append(itemSiteData, ItemSiteData{
		SiteId:     1,
		Name:       "htmlHandler1",
		ParserType: htmlParser,
		Path:       []interface{}{".check"},
		Url:        siteToUrl["htmlHandler1"]})

	itemSiteData = append(itemSiteData, ItemSiteData{
		SiteId:     2,
		Name:       "htmlHandler2",
		ParserType: htmlParser,
		Path:       []interface{}{".check1"},
		Url:        siteToUrl["htmlHandler2"]})

	itemSiteData = append(itemSiteData, ItemSiteData{
		SiteId:     3,
		Name:       "jsonHandler1",
		ParserType: jsonParser,
		Path:       []interface{}{"price"},
		Url:        siteToUrl["jsonHandler1"]})

	itemSiteData = append(itemSiteData, ItemSiteData{
		SiteId:     4,
		Name:       "jsonHandler2",
		ParserType: jsonParser,
		Path:       []interface{}{".check"},
		Url:        siteToUrl["jsonHandler2"]})

	itemData := ItemData{
		Name:        "Box",
		ItemId:      1,
		SiteData:    itemSiteData,
		SiteHistory: make(map[string]*SiteHistory)}

	itemData.SiteHistory["htmlHandler1"] = &SiteHistory{}
	itemData.SiteHistory["htmlHandler2"] = &SiteHistory{}
	itemData.SiteHistory["jsonHandler1"] = &SiteHistory{}
	itemData.SiteHistory["jsonHandler2"] = &SiteHistory{}

	return itemData
}
