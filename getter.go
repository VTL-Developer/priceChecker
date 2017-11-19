package main

func GetAndUpdateDataForItems(itemsToUpdate []*ItemData, numOfThreads int) []*ItemData {
	var nextItem *ItemData = nil

	if numOfThreads < 0 {
		panic("Number of threads should be 1 or greater")
	}

	idChan := make(chan int)
	jobChans := startThreads(itemsToUpdate, numOfThreads, idChan)
	numOfJobs := len(jobChans)
	leftover := itemsToUpdate[min(numOfJobs, len(itemsToUpdate)):]

	var threadId int
	logDebug("Have %v active thread(s) running", numOfJobs)
	for numOfJobs > 0 {
		threadId = <-idChan
		nextItem, leftover = getNextData(leftover)
		jobChans[threadId] <- nextItem

		if nextItem == nil {
			numOfJobs -= 1
			close(jobChans[threadId])
			logDebug("A thread has completed its job, %v remaining active thread(s)", numOfJobs)
		}
	}

	close(idChan)
	return itemsToUpdate
}

func getNextData(itemsToWorkOn []*ItemData) (*ItemData, []*ItemData) {
	if len(itemsToWorkOn) > 0 {
		return itemsToWorkOn[0], itemsToWorkOn[1:]
	}

	return nil, itemsToWorkOn
}

func getAndUpdateDataForItemThread(currentItem *ItemData, id int, jobChan chan *ItemData, idChan chan int) {
	for currentItem != nil {
		getAndUpdateData(currentItem)
		idChan <- id
		currentItem = <-jobChan
		logDebug("Received job: %v", currentItem)
	}

	logDebug("Closing the channel")
}

func getAndUpdateData(item *ItemData) {
	logInfo(`Processing price update for %v`, item.Name)

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
				logInfo("Found the cost of %v from %v for %v", price, siteData.Name, item.Name)
			}
		}
	}

	logInfo("Finished processing %v\n%q", item.Name, item)
}

func startSingleThread(item *ItemData, id int, idChan chan int) chan *ItemData {
	jobChan := make(chan *ItemData)
	go getAndUpdateDataForItemThread(item, id, jobChan, idChan)
	return jobChan
}

func startThreads(allItems []*ItemData, numOfThreads int, idChan chan int) []chan *ItemData {
	length := len(allItems)
	var idChans []chan *ItemData

	for threadId := 0; threadId < length && threadId < numOfThreads; threadId++ {
		idChans = append(idChans, startSingleThread(allItems[threadId], threadId, idChan))
	}

	return idChans
}
