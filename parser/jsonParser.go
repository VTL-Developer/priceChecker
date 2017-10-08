package main

import (
//"encoding/json"
)

func GetCost(content interface{}, contentPath []interface{}) float64 {
	jsonObj := content
	cost := -1.0
	found := true
	converted := true
	name := ""
	index := -1
	for _, path := range contentPath {
		index, converted = path.(int)

		if converted {
			jsonObj, found = getJsonObjectByIndex(jsonObj, index)
		} else {
			name, converted = path.(string)
			if converted {
				jsonObj, found = getJsonObjectByName(jsonObj, name)
			}
		}

		if !converted || !found {
			return -1.0
		}
	}

	if found {
		cost, converted = jsonObj.(float64)
	}

	if !converted {
		cost = -1.0
	}

	return cost
}

func getJsonObjectByName(content interface{}, name string) (interface{}, bool) {
	jsonObj, ok := content.(map[string]interface{})

	if !ok {
		return nil, ok
	}

	innerObj, ok := jsonObj[name].(interface{})
	return innerObj, ok
}

func getJsonObjectByIndex(content interface{}, index int) (interface{}, bool) {
	jsonObj, ok := content.([]interface{})

	if !ok || len(jsonObj) < index+1 {
		return nil, false
	}

	innerObj, ok := jsonObj[index].(interface{})
	return innerObj, ok
}

func main() {
	// content := make(map[string] interface{})
	// e := json.Unmarshal(data, &content)

	// if e != nil {
	// 	panic(e)
	// }
}
