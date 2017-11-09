package main

import (
	"testing"
	"time"
)

func TestAddPriceToLast24Hour_EmptyCase(t *testing.T) {
	s := SiteHistory{}
	s.AddPriceToLast24HourHistory(50.00)
	validateSiteHistoryList(
		t,
		&s,
		[]PriceHistory{
			makeMockPriceHistories(50.00, time.Now().UTC()),
		},
		[]PriceHistory{})
}

func TestAddPriceToLast24Hour_WhenTheresData(t *testing.T) {
	datetime := make([]time.Time, 3)
	datetime[0] = time.Now().UTC()
	datetime[1] = time.Now().UTC()

	s := SiteHistory{
		Last24HourPriceHistory: []PriceHistory{
			makeMockPriceHistories(50.0, datetime[0]),
			makeMockPriceHistories(120.0, datetime[1]),
		},
	}

	s.AddPriceToLast24HourHistory(102.5)
	datetime[2] = time.Now().UTC()
	validateSiteHistoryList(
		t,
		&s,
		[]PriceHistory{
			makeMockPriceHistories(50.00, datetime[0]),
			makeMockPriceHistories(120.00, datetime[1]),
			makeMockPriceHistories(102.50, datetime[2]),
		},
		[]PriceHistory{})
}

func TestAddPriceToLast24Hour_MultipleCase(t *testing.T) {
	s := SiteHistory{}
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
		[]PriceHistory{
			makeMockPriceHistories(50.00, datetime[0]),
			makeMockPriceHistories(120.00, datetime[1]),
			makeMockPriceHistories(102.50, datetime[2]),
		},
		[]PriceHistory{})
}

func TestAddPriceToDayByDayHistory_EmptyCase(t *testing.T) {
	s := SiteHistory{}
	s.AddPriceToDayByDayHistory(50.00)
	validateSiteHistoryList(
		t,
		&s,
		[]PriceHistory{},
		[]PriceHistory{
			makeMockPriceHistories(50.00, time.Now().UTC()),
		})
}

func TestAddPriceToDayByDayHistory_WhenTheresData(t *testing.T) {
	datetime := make([]time.Time, 3)
	datetime[0] = time.Now().UTC()
	datetime[1] = time.Now().UTC()

	s := SiteHistory{
		DayByDayPriceHistory: []PriceHistory{
			makeMockPriceHistories(50.0, datetime[0]),
			makeMockPriceHistories(120.0, datetime[1]),
		},
	}

	s.AddPriceToDayByDayHistory(102.5)
	datetime[2] = time.Now().UTC()
	validateSiteHistoryList(
		t,
		&s,
		[]PriceHistory{},
		[]PriceHistory{
			makeMockPriceHistories(50.00, datetime[0]),
			makeMockPriceHistories(120.00, datetime[1]),
			makeMockPriceHistories(102.50, datetime[2]),
		})
}

func TestAddPriceToDayByDayHistory_MultipleCase(t *testing.T) {
	s := SiteHistory{}
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
		[]PriceHistory{},
		[]PriceHistory{
			makeMockPriceHistories(50.00, datetime[0]),
			makeMockPriceHistories(120.00, datetime[1]),
			makeMockPriceHistories(102.50, datetime[2]),
		})
}

func TestTrimLast24HoursToN_ToLast1Item(t *testing.T) {
	datetimes := getMockSetOfTime(9)
	mockPriceHistories := make([]PriceHistory, 9)

	for i, datetime := range datetimes {
		mockPriceHistories[i] = makeMockPriceHistories(1.0, datetime)
	}

	s := SiteHistory{
		Last24HourPriceHistory: mockPriceHistories,
	}

	s.TrimLast24HoursToN(1)

	validateSiteHistoryList(
		t,
		&s,
		mockPriceHistories[8:],
		[]PriceHistory{})
}

func TestTrimLast24HoursToN_ToLast5Item(t *testing.T) {
	datetimes := getMockSetOfTime(9)
	mockPriceHistories := make([]PriceHistory, 9)

	for i, datetime := range datetimes {
		mockPriceHistories[i] = makeMockPriceHistories(1.0, datetime)
	}

	s := SiteHistory{
		Last24HourPriceHistory: mockPriceHistories,
	}

	s.TrimLast24HoursToN(5)

	validateSiteHistoryList(
		t,
		&s,
		mockPriceHistories[4:],
		[]PriceHistory{})
}

func TestTrimLast24HoursToN_ToLastNGreaterThanCount(t *testing.T) {
	datetimes := getMockSetOfTime(9)
	mockPriceHistories := make([]PriceHistory, 9)

	for i, datetime := range datetimes {
		mockPriceHistories[i] = makeMockPriceHistories(1.0, datetime)
	}

	s := SiteHistory{
		Last24HourPriceHistory: mockPriceHistories,
	}

	s.TrimLast24HoursToN(900)

	validateSiteHistoryList(
		t,
		&s,
		mockPriceHistories[:],
		[]PriceHistory{})
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

func makeMockPriceHistories(price float64, datetime time.Time) PriceHistory {
	return PriceHistory{
		Price:    price,
		Datetime: datetime,
	}
}

func validateSiteHistoryList(t *testing.T, s *SiteHistory, expectedLast24PriceHistory []PriceHistory, expectedDayToDayHistory []PriceHistory) {
	validateSiteHistory(t, "Last24HourPriceHistory", s.Last24HourPriceHistory, expectedLast24PriceHistory)
	validateSiteHistory(t, "DayByDayPriceHistory", s.DayByDayPriceHistory, expectedDayToDayHistory)
}

func validateSiteHistory(t *testing.T, typeOfHistory string, actualPriceHistory []PriceHistory, expectedPriceHistory []PriceHistory) {
	if len(actualPriceHistory) != len(expectedPriceHistory) {
		t.Errorf("length of %v and expected not the same, %v and %v, respectively", typeOfHistory,
			len(actualPriceHistory), len(expectedPriceHistory))
	}

	for i, ph := range actualPriceHistory {
		if ph.Price != expectedPriceHistory[i].Price {
			t.Errorf("price at [%v] is %v, should be %v", i,
				ph.Price, expectedPriceHistory[i].Price)
		}

		timeDiff := expectedPriceHistory[i].Datetime.Sub(ph.Datetime)

		if timeDiff.Seconds() > 1 {
			t.Errorf("datetime at [%v] is %v, should be closer to %v", i,
				ph.Datetime, expectedPriceHistory[i].Datetime)
		}
	}
}
