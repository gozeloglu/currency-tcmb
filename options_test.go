package currency

import (
	"testing"
	"time"
)

func TestWithDate(t *testing.T) {
	testCases := []struct {
		name    string
		day     int
		month   time.Month
		year    int
		expDate string
	}{
		{
			name:    "D M YYYY format",
			day:     1,
			month:   time.March,
			year:    2023,
			expDate: "01032023",
		},
		{
			name:    "DD M YYYY format",
			day:     12,
			month:   time.May,
			year:    2023,
			expDate: "12052023",
		},
		{
			name:    "D MM YYYY format",
			day:     3,
			month:   time.December,
			year:    2023,
			expDate: "03122023",
		},
		{
			name:    "DD MM YYYY format",
			day:     20,
			month:   time.November,
			year:    2023,
			expDate: "20112023",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tcmb := &TCMB{}
			o := WithDate(tc.day, tc.month, tc.year)
			o(tcmb)

			if tcmb.date != tc.expDate {
				t.Errorf("expected date: %s\nactual date: %s", tc.expDate, tcmb.date)
			}
			t.Logf("date: %s", tcmb.date)
		})
	}
}
