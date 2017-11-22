package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetCostJson_EmptyCase(t *testing.T) {
	testGetCostJson(t, []byte(`{}`), []interface{}{}, -1.0)
}

func TestGetCostJson_EmptyCaseWithPath(t *testing.T) {
	testGetCostJson(t, []byte(`{}`), []interface{}{"Somewhere"}, -1.0)
}

func TestGetCostJson_SimpleCase(t *testing.T) {
	testGetCostJson(t, []byte(`{"cost": 123.45}`), []interface{}{"cost"}, 123.45)
}

func TestGetCostJson_SimpleCaseWithMiss(t *testing.T) {
	testGetCostJson(t, []byte(`{"cost": 123.45}`), []interface{}{"price"}, -1)
}

func TestGetCostJson_SimpleCaseWithCaseSensitivityMiss(t *testing.T) {
	testGetCostJson(t, []byte(`{"cost": 123.45}`), []interface{}{"Cost"}, -1)
}

func TestGetCostJson_ComplexJsonObject(t *testing.T) {
	testGetCostJson(t, []byte(`{
		"item": "name",
		"detail": {
			"item": {
				"cost": 123.45
			},
			"cost": 2
		},
		"cost": 4
	}`), []interface{}{"detail", "item", "cost"}, 123.45)
}

func TestGetCostJson_ComplexJsonObjectWithIntsAsNamedPaths(t *testing.T) {
	testGetCostJson(t, []byte(`{
		"item": "name",
		"detail": {
			"52": {
				"cost": 123.45
			},
			"cost": 2
		},
		"cost": 4
	}`), []interface{}{"detail", "52", "cost"}, 123.45)
}

func TestGetCostJson_ComplexJsonObjectWithIntsAsNamedPathsAndArrays(t *testing.T) {
	testGetCostJson(t, []byte(`{
		"item":"name",
		"detail": [{
			"52": {
				"cost": [null, null, 123.45, null]
			},
			"cost": 2
		}],
		"cost": 4
	}`), []interface{}{"detail", 0, "52", "cost", 2}, 123.45)
}

func Test_getDataAsJson_With_PriceHistoryJson(t *testing.T) {
	var rsp http.Response
	rsp.Body = &MockClosingBuffer{bytes.NewBufferString(`{"price":"54.05", "datetime":"2017-10-24T22:15:21.000000Z"}`)}
	parser := JsonParser{}
	jsonBody, err := parser.GetData(&rsp)

	AssertNotError(err, t)

	if convertedJsonBody, ok := (*jsonBody).(map[string]interface{}); !ok {
		t.Errorf("Conversion was not successful.\nOriginal: %q\nConverted: %q", jsonBody, convertedJsonBody)
	} else {
		AssertEqual(convertedJsonBody["price"], "54.05", t)
		AssertEqual(convertedJsonBody["datetime"], "2017-10-24T22:15:21.000000Z", t)
	}
}

func Test_getDataAsJson_With_BadJSON(t *testing.T) {
	var rsp http.Response
	rsp.Body = &MockClosingBuffer{bytes.NewBufferString(`{"price":"54.05", "datetime":"2017-10-24T22:15:21.000000Z"`)}
	parser := JsonParser{}
	_, err := parser.GetData(&rsp)

	AssertError(err, t)
}

func testGetCostJson(t *testing.T, jsonBody []byte, path []interface{}, expected float64) {
	var jsonObj interface{}
	err := json.Unmarshal(jsonBody, &jsonObj)

	AssertNotError(err, t)

	parser := JsonParser{}
	got := parser.GetCost(&jsonObj, path)
	AssertEqual(got, expected, t)
}
