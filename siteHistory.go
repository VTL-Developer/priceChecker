package main

import (
	"time"
)

type SiteHistory struct {
	Last24HourPriceHistory []PriceHistory
	DayByDayPriceHistory   []PriceHistory
}

func (s *SiteHistory) AddPriceToLast24HourHistory(price float64) {
	s.Last24HourPriceHistory = addPriceToPriceHistory(price, s.Last24HourPriceHistory)
}

func (s *SiteHistory) AddPriceToDayByDayHistory(price float64) {
	s.DayByDayPriceHistory = addPriceToPriceHistory(price, s.DayByDayPriceHistory)
}

func (s *SiteHistory) TrimLast24HoursToN(n int) {
	length := len(s.Last24HourPriceHistory)

	if length > n {
		s.Last24HourPriceHistory = s.Last24HourPriceHistory[length-n:]
	}
}

func addPriceToPriceHistory(price float64, priceHistorySlice []PriceHistory) []PriceHistory {
	if price > 0 {
		priceHistorySlice = (append(
			priceHistorySlice,
			PriceHistory{
				Price:    price,
				Datetime: time.Now().UTC()}))
	}

	return priceHistorySlice
}
