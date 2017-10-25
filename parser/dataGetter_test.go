package main

import (
	"bytes"
	"net/http"
	"testing"
)

type MockClosingBuffer struct {
	*bytes.Buffer
}

func (mcb *MockClosingBuffer) Close() error {
	return nil
}

func Test_getDataAsJson_With_PriceHistory(t *testing.T) {
	var rsp http.Response
	var jsonBody interface{}
	rsp.Body = &MockClosingBuffer{bytes.NewBufferString(`{"price":"54.05", "datetime":"2017-10-24T22:15:21.000000Z"}`)}
	err := getDataAsJson(&rsp, &jsonBody)

	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	convertedJsonBody, ok := jsonBody.(map[string]interface{})

	if !ok {
		t.Errorf("Conversion was not successful.\nOriginal: %q\nConverted: %q", jsonBody, convertedJsonBody)
	}

	if convertedJsonBody["price"] != "54.05" {
		t.Errorf("JSON body price should is %q, should be %q", convertedJsonBody["price"], "54.05")
	}

	if convertedJsonBody["datetime"] != "2017-10-24T22:15:21.000000Z" {
		t.Errorf("JSON body datetime should is %q, should be %q", convertedJsonBody["datetime"], "2017-10-24T22:15:21.000000Z")
	}
}
