package main

import (
	"time"
)

type siteHistory struct {
	Last24HourPriceHistory []priceHistory
	DayByDayPriceHistory   []priceHistory
}

func (s *siteHistory) AddPriceToLast24HourHistory(price float64) {
	s.Last24HourPriceHistory = addPriceToPriceHistory(price, s.Last24HourPriceHistory)
}

func (s *siteHistory) AddPriceToDayByDayHistory(price float64) {
	s.DayByDayPriceHistory = addPriceToPriceHistory(price, s.DayByDayPriceHistory)
}

func (s *siteHistory) TrimLast24HoursToN(n int) {
	length := len(s.Last24HourPriceHistory)

	if length > n {
		s.Last24HourPriceHistory = s.Last24HourPriceHistory[length-n:]
	}
}

func addPriceToPriceHistory(price float64, priceHistorySlice []priceHistory) []priceHistory {
	if price > 0 {
		priceHistorySlice = (append(
			priceHistorySlice,
			priceHistory{
				price:    price,
				datetime: time.Now().UTC()}))
	}

	return priceHistorySlice
}
