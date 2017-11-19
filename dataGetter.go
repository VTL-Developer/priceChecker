package main

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"time"
)

var client = http.Client{Timeout: 10 * time.Second}

func getDataAsJson(rsp *http.Response) (*interface{}, error) {
	var target interface{}
	defer closeRsp(rsp)
	jsonBody, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		logError("Error trying to transform the HTTP response to JSON object for %v , \nException: %v",
			rsp.Request.URL.String(), err)
		return nil, err
	}

	err = json.Unmarshal(jsonBody, &target)
	if err != nil {
		logError("Issue transforming JSON: %v", err)
		logDebug("Content was:\n%v", string(jsonBody))
	}

	return &target, err
}

func getDataAsHtml(rsp *http.Response) (*interface{}, error) {
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

func getHttpBodyResponse(url string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	if headers != nil {
		for header, value := range headers {
			req.Header.Set(header, value)
		}
	}

	logDebug("Request made: %q", req)
	return client.Do(req)
}

func closeRsp(rsp *http.Response) {
	if rsp.Request != nil {
		logInfo("Closing the connection for %v", rsp.Request.URL)
	}

	if err := rsp.Body.Close(); err != nil {
		logError("Error closing response: %v", err)
	}

	if rsp.Request != nil {
		logInfo("Closed the connection for %v", rsp.Request.URL)
	}
}
