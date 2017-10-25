package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

var client = http.Client{Timeout: 10 * time.Second}

func getDataAsJson(rsp *http.Response, target interface{}) error {
	defer rsp.Body.Close()
	jsonBody, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonBody, target)
	return err
}

func getHttpBodyResponse(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	if headers != nil {
		for header, value := range headers {
			req.Header.Add(header, value)
		}
	}

	rsp, err := client.Do(req)
	defer rsp.Body.Close()

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(rsp.Body)
}
