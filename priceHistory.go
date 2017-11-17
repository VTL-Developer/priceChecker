package main

import (
	"time"
)

type PriceHistory struct {
	Price    float64
	Datetime time.Time
}

func MakePriceEntry(price float64) PriceHistory {
	return PriceHistory{
		Price:    price,
		Datetime: time.Now().UTC()}
}
