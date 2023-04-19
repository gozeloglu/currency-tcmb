package currency

import (
	"fmt"
	"time"
)

type OptionFunc func(t *TCMB)

// WithDate enables to pass custom date for past dates.
func WithDate(day int, month time.Month, year int) OptionFunc {
	return func(t *TCMB) {
		d := fmt.Sprintf("%d", day)
		m := fmt.Sprintf("%d", month)
		if day < 10 {
			d = "0" + d
		}
		if month < 10 {
			m = "0" + m
		}
		date := fmt.Sprintf("%s%s%d", d, m, year)
		t.date = date
	}
}
