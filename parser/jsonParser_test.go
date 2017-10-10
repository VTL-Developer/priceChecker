package main

import (
	"encoding/json"
	"testing"
)

func TestGetCost_EmptyCase(t *testing.T) {
	testGetCost(t, []byte(`{}`), []interface{}{}, -1.0)
}

func TestGetCost_EmptyCaseWithPath(t *testing.T) {
	testGetCost(t, []byte(`{}`), []interface{}{"Somewhere"}, -1.0)
}

func TestGetCost_SimpleCase(t *testing.T) {
	testGetCost(t, []byte(`{"cost": 123.45}`), []interface{}{"cost"}, 123.45)
}

func TestGetCost_SimpleCaseWithMiss(t *testing.T) {
	testGetCost(t, []byte(`{"cost": 123.45}`), []interface{}{"price"}, -1)
}

func TestGetCost_SimpleCaseWithCaseSensitivityMiss(t *testing.T) {
	testGetCost(t, []byte(`{"cost": 123.45}`), []interface{}{"Cost"}, -1)
}

func TestGetCost_ComplexJsonObject(t *testing.T) {
	testGetCost(t, []byte(`{
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

func TestGetCost_ComplexJsonObjectWithIntsAsNamedPaths(t *testing.T) {
	testGetCost(t, []byte(`{
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

func TestGetCost_ComplexJsonObjectWithIntsAsNamedPathsAndArrays(t *testing.T) {
	testGetCost(t, []byte(`{
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

func testGetCost(t *testing.T, jsonBody []byte, path []interface{}, expected float64) {
	var jsonObj interface{}
	e := json.Unmarshal(jsonBody, &jsonObj)

	if e != nil {
		t.Errorf("Error thrown: %v", e)
		return
	}

	got := GetCost(&jsonObj, &path)
	if got != expected {
		t.Errorf("GetCost returned %q, should be %q", got, expected)
	}
}
