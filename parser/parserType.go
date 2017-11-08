package main

import (
	"net/http"
)

type ParserType int

const (
	unknownParserType ParserType = iota
	jsonParser        ParserType = iota
	htmlParser        ParserType = iota
)

func getDataTransformer(p ParserType) func(rsp *http.Response) (*interface{}, error) {
	switch p {
	case jsonParser:
		return getDataAsJson
	case htmlParser:
		return getDataAsHtml
	default:
		return nil
	}
}

func getDataParser(p ParserType) func(content *interface{}, contentPath []interface{}) float64 {
	switch p {
	case jsonParser:
		return GetCostJson
	case htmlParser:
		return GetCostHtml
	default:
		return nil
	}
}
