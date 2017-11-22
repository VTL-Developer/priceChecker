package main

import (
	"net/http"
)

type Parser interface {
	GetCost(content *interface{}, contentPath []interface{}) float64
	GetData(rsp *http.Response) (*interface{}, error)
}

func GetParser(p ParserType) Parser {
	switch p {
	case jsonParser:
		return JsonParser{}
	case htmlParser:
		return HtmlParser{}
	default:
		return nil
	}
}
