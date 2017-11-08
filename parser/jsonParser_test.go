package main

import (
	"encoding/json"
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

func testGetCostJson(t *testing.T, jsonBody []byte, path []interface{}, expected float64) {
	var jsonObj interface{}
	e := json.Unmarshal(jsonBody, &jsonObj)

	if e != nil {
		t.Errorf("Error thrown: %v", e)
		return
	}

	got := GetCostJson(&jsonObj, path)
	if got != expected {
		t.Errorf("GetCostJson returned %q, should be %q", got, expected)
	}
}
