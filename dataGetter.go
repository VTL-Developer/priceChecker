package main

import (
	"net/http"
	"time"
)

var client = http.Client{Timeout: 10 * time.Second}

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
