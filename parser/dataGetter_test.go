package main

import (
	"testing"
)

func Test_getDataAsJson_With_PriceHistory(t *testing.T) {
	var priceHistory PriceHistory
	getDataAsJson([]byte(`{"price":54.05, "datetime":"2017-10-24T22:15:21.000000Z"}`), &priceHistory)

	if priceHistory.Price != 54.05 {
		t.Errorf("Price is %q, but should be %q", priceHistory.Price, 54.05)
	}

	_, zone := priceHistory.Datetime.Zone()

	if zone != 0 {
		t.Errorf("Datetime Zone is %q, but should be %q", zone, 0)
	}

	if priceHistory.Datetime.Hour() != 22 {
		t.Errorf("Datetime Hour is %q, but should be %q", priceHistory.Datetime.Hour(), 22)
	}

	if priceHistory.Datetime.Minute() != 15 {
		t.Errorf("Datetime Minute is %q, but should be %q", priceHistory.Datetime.Minute(), 15)
	}

	if priceHistory.Datetime.Second() != 21 {
		t.Errorf("Datetime Second is %q, but should be %q", priceHistory.Datetime.Second(), 21)
	}

	if priceHistory.Datetime.Year() != 2017 {
		t.Errorf("Datetime Year is %q, but should be %q", priceHistory.Datetime.Year(), 2017)
	}

	if priceHistory.Datetime.Month() != 10 {
		t.Errorf("Datetime Month is %q, but should be %q", priceHistory.Datetime.Month(), 10)
	}

	if priceHistory.Datetime.Day() != 24 {
		t.Errorf("Datetime Day is %q, but should be %q", priceHistory.Datetime.Day(), 24)
	}
}
