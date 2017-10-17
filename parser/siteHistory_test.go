package main

import (
	"testing"
	"time"
)

func TestAddPriceToLast24Hour_EmptyCase(t *testing.T) {
	s := siteHistory{}
	s.AddPriceToLast24HourHistory(50.00)
	validateSiteHistoryList(
		t,
		&s,
		[]priceHistory{
			makeMockPriceHistories(50.00, time.Now().UTC()),
		},
		[]priceHistory{})
}

func TestAddPriceToLast24Hour_WhenTheresData(t *testing.T) {
	datetime := make([]time.Time, 3)
	datetime[0] = time.Now().UTC()
	datetime[1] = time.Now().UTC()

	s := siteHistory{
		Last24HourPriceHistory: []priceHistory{
			makeMockPriceHistories(50.0, datetime[0]),
			makeMockPriceHistories(120.0, datetime[1]),
		},
	}

	s.AddPriceToLast24HourHistory(102.5)
	datetime[2] = time.Now().UTC()
	validateSiteHistoryList(
		t,
		&s,
		[]priceHistory{
			makeMockPriceHistories(50.00, datetime[0]),
			makeMockPriceHistories(120.00, datetime[1]),
			makeMockPriceHistories(102.50, datetime[2]),
		},
		[]priceHistory{})
}

func TestAddPriceToLast24Hour_MultipleCase(t *testing.T) {
	s := siteHistory{}
	datetime := make([]time.Time, 3)
	s.AddPriceToLast24HourHistory(50.00)
	datetime[0] = time.Now().UTC()
	s.AddPriceToLast24HourHistory(120)
	datetime[1] = time.Now().UTC()
	s.AddPriceToLast24HourHistory(102.5)
	datetime[2] = time.Now().UTC()

	validateSiteHistoryList(
		t,
		&s,
		[]priceHistory{
			makeMockPriceHistories(50.00, datetime[0]),
			makeMockPriceHistories(120.00, datetime[1]),
			makeMockPriceHistories(102.50, datetime[2]),
		},
		[]priceHistory{})
}

func TestAddPriceToDayByDayHistory_EmptyCase(t *testing.T) {
	s := siteHistory{}
	s.AddPriceToDayByDayHistory(50.00)
	validateSiteHistoryList(
		t,
		&s,
		[]priceHistory{},
		[]priceHistory{
			makeMockPriceHistories(50.00, time.Now().UTC()),
		})
}

func TestAddPriceToDayByDayHistory_WhenTheresData(t *testing.T) {
	datetime := make([]time.Time, 3)
	datetime[0] = time.Now().UTC()
	datetime[1] = time.Now().UTC()

	s := siteHistory{
		DayByDayPriceHistory: []priceHistory{
			makeMockPriceHistories(50.0, datetime[0]),
			makeMockPriceHistories(120.0, datetime[1]),
		},
	}

	s.AddPriceToDayByDayHistory(102.5)
	datetime[2] = time.Now().UTC()
	validateSiteHistoryList(
		t,
		&s,
		[]priceHistory{},
		[]priceHistory{
			makeMockPriceHistories(50.00, datetime[0]),
			makeMockPriceHistories(120.00, datetime[1]),
			makeMockPriceHistories(102.50, datetime[2]),
		})
}

func TestAddPriceToDayByDayHistory_MultipleCase(t *testing.T) {
	s := siteHistory{}
	datetime := make([]time.Time, 3)
	s.AddPriceToDayByDayHistory(50.00)
	datetime[0] = time.Now().UTC()
	s.AddPriceToDayByDayHistory(120)
	datetime[1] = time.Now().UTC()
	s.AddPriceToDayByDayHistory(102.5)
	datetime[2] = time.Now().UTC()

	validateSiteHistoryList(
		t,
		&s,
		[]priceHistory{},
		[]priceHistory{
			makeMockPriceHistories(50.00, datetime[0]),
			makeMockPriceHistories(120.00, datetime[1]),
			makeMockPriceHistories(102.50, datetime[2]),
		})
}

func TestTrimLast24HoursToN_ToLast1Item(t *testing.T) {
	datetimes := getMockSetOfTime(9)
	mockPriceHistories := make([]priceHistory, 9)

	for i, datetime := range datetimes {
		mockPriceHistories[i] = makeMockPriceHistories(1.0, datetime)
	}

	s := siteHistory{
		Last24HourPriceHistory: mockPriceHistories,
	}

	s.TrimLast24HoursToN(1)

	validateSiteHistoryList(
		t,
		&s,
		mockPriceHistories[8:],
		[]priceHistory{})
}

func TestTrimLast24HoursToN_ToLast5Item(t *testing.T) {
	datetimes := getMockSetOfTime(9)
	mockPriceHistories := make([]priceHistory, 9)

	for i, datetime := range datetimes {
		mockPriceHistories[i] = makeMockPriceHistories(1.0, datetime)
	}

	s := siteHistory{
		Last24HourPriceHistory: mockPriceHistories,
	}

	s.TrimLast24HoursToN(5)

	validateSiteHistoryList(
		t,
		&s,
		mockPriceHistories[4:],
		[]priceHistory{})
}

func TestTrimLast24HoursToN_ToLastNGreaterThanCount(t *testing.T) {
	datetimes := getMockSetOfTime(9)
	mockPriceHistories := make([]priceHistory, 9)

	for i, datetime := range datetimes {
		mockPriceHistories[i] = makeMockPriceHistories(1.0, datetime)
	}

	s := siteHistory{
		Last24HourPriceHistory: mockPriceHistories,
	}

	s.TrimLast24HoursToN(900)

	validateSiteHistoryList(
		t,
		&s,
		mockPriceHistories[:],
		[]priceHistory{})
}

func getMockSetOfTime(count int) []time.Time {
	datetimes := make([]time.Time, count)
	now := time.Now()
	for i := 0; i < count; i++ {
		datetimes[i] = now
		now = now.Add(time.Minute)
	}

	return datetimes
}

func makeMockPriceHistories(price float64, datetime time.Time) priceHistory {
	return priceHistory{
		price:    price,
		datetime: datetime,
	}
}

func validateSiteHistoryList(t *testing.T, s *siteHistory, expectedLast24PriceHistory []priceHistory, expectedDayToDayHistory []priceHistory) {
	validateSiteHistory(t, "Last24HourPriceHistory", s.Last24HourPriceHistory, expectedLast24PriceHistory)
	validateSiteHistory(t, "DayByDayPriceHistory", s.DayByDayPriceHistory, expectedDayToDayHistory)
}

func validateSiteHistory(t *testing.T, typeOfHistory string, actualPriceHistory []priceHistory, expectedPriceHistory []priceHistory) {
	if len(actualPriceHistory) != len(expectedPriceHistory) {
		t.Errorf("length of %v and expected not the same, %v and %v, respectively", typeOfHistory,
			len(actualPriceHistory), len(expectedPriceHistory))
	}

	for i, ph := range actualPriceHistory {
		if ph.price != expectedPriceHistory[i].price {
			t.Errorf("price at [%v] is %v, should be %v", i,
				ph.price, expectedPriceHistory[i].price)
		}

		timeDiff := expectedPriceHistory[i].datetime.Sub(ph.datetime)

		if timeDiff.Seconds() > 1 {
			t.Errorf("datetime at [%v] is %v, should be closer to %v", i,
				ph.datetime, expectedPriceHistory[i].datetime)
		}
	}
}
