package main

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
)

type HtmlParser struct {
}

func (h HtmlParser) GetCost(content *interface{}, contentPath []interface{}) float64 {
	var queryPath string

	document, converted := (*content).(goquery.Document)

	if !converted {
		logError("Input was not a goquery document object: %q", content)
		return -1.0
	}

	node := document.Find("body")

	for _, item := range contentPath {
		queryPath, converted = item.(string)

		if converted {
			node = node.Find(queryPath)
		}

		if !converted || node == nil {
			logWarning("Was not able to find the cost from HTML page.")
			return -1.0
		}
	}

	text := node.Text()
	cost, err := strconv.ParseFloat(text, 64)

	if err != nil {
		logWarning("Unable to convert the text to float: %v", err)
		return -1.0
	}

	logInfo("Found the cost: %v", cost)
	return cost
}

func (h HtmlParser) GetData(rsp *http.Response) (*interface{}, error) {
	var convertedDocument interface{}
	var err error
	defer closeRsp(rsp)

	document, err := goquery.NewDocumentFromResponse(rsp)

	if err != nil {
		logError("Error trying to transform the HTTP response to GoQuery object for %v , \nException: %v",
			rsp.Request.URL.String(), err)
		return nil, err
	}

	convertedDocument = *document
	return &convertedDocument, nil
}
