package main

type ItemData struct {
	Name        string
	ItemId      int
	SiteData    []ItemSiteData
	SiteHistory map[string]*SiteHistory
}
