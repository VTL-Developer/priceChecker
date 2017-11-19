package main

func GetCostJson(content *interface{}, contentPath []interface{}) float64 {
	jsonObj := content
	cost := -1.0
	found := true
	converted := true
	name := ""
	index := -1

	logDebug("GetCost function\nJSON body is: %v\nJSON path to check: %v", *content, contentPath)

	for _, path := range contentPath {
		if index, converted = path.(int); converted {
			jsonObj, found = getJsonObjectByIndex(jsonObj, index)
		} else {
			if name, converted = path.(string); converted {
				jsonObj, found = getJsonObjectByName(jsonObj, name)
			}
		}

		if !converted || !found {
			logWarning("Was not able to find the cost from JSON.")
			return -1.0
		}
	}

	if found {
		cost, converted = (*jsonObj).(float64)
	}

	if !converted {
		cost = -1.0
	}

	logInfo("Found the cost: %v", cost)
	return cost
}

func getJsonObjectByName(content *interface{}, name string) (*interface{}, bool) {
	if jsonObj, ok := (*content).(map[string]interface{}); !ok {
		return nil, ok
	} else {
		innerObj, ok := jsonObj[name].(interface{})
		return &innerObj, ok
	}
}

func getJsonObjectByIndex(content *interface{}, index int) (*interface{}, bool) {
	if jsonObj, ok := (*content).([]interface{}); !ok || len(jsonObj) < index+1 {
		return nil, false
	} else {
		innerObj, ok := jsonObj[index].(interface{})
		return &innerObj, ok
	}
}
