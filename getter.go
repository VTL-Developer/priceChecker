package main

func getAndUpdateData(item *ItemData) {
	logDebug(`Processing "%v"`, item.Name)

	for _, siteData := range item.SiteData {
		var content *interface{}
		logDebug(`Parsing "%v" for "%v"`, siteData.Name, item.Name)

		dataTransformer := getDataTransformer(siteData.ParserType)
		dataParser := getDataParser(siteData.ParserType)

		if dataTransformer == nil || dataParser == nil {
			logWarning(`Skip "%v" for "%v"`, siteData.Name, item.Name)
		}

		body, err := getHttpBodyResponse(siteData.Url, nil)
		if err != nil {
			logError("Issue with HTTP Request: %v", err)
			continue
		}

		content, err = dataTransformer(body)
		if err != nil {
			logError("Issue with data transformation: %v", err)
			continue
		}

		price := dataParser(content, siteData.Path)
		if price > 0 {
			if history, ok := item.SiteHistory[siteData.Name]; ok {
				history.AddPriceToLast24HourHistory(price)
				logDebug("History: %q", history)
			}
		}
	}

	logDebug(`Finished processing "%v"\n%q`, item.Name, item)
}
