package main

type ItemSiteData struct {
	Name       string
	SiteId     int
	Url        string
	ParserType ParserType
	Path       []interface{}
}
