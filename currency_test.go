package currency

import "testing"

func TestTCMB_FromCurrencyCode(t *testing.T) {
	tcmb := &TCMB{
		forexBuying: map[Code]string{
			USD: "10",
			EUR: "15",
		},
		forexSelling: map[Code]string{
			USD: "11",
			EUR: "16",
		},
		banknoteBuying: map[Code]string{
			USD: "12",
			EUR: "17",
		},
		banknoteSelling: map[Code]string{
			USD: "13",
			EUR: "18",
		},
		date: "20231904",
	}

	testCases := []struct {
		name            string
		code            Code
		forexBuying     string
		forexSelling    string
		banknoteBuying  string
		banknoteSelling string
	}{
		{
			name:            "USD",
			code:            USD,
			forexBuying:     "10",
			forexSelling:    "11",
			banknoteBuying:  "12",
			banknoteSelling: "13",
		},
		{
			name:            "EUR",
			code:            EUR,
			forexBuying:     "15",
			forexSelling:    "16",
			banknoteBuying:  "17",
			banknoteSelling: "18",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			curr := tcmb.FromCurrencyCode(tc.code)
			fs := curr.ForexSelling()
			fb := curr.ForexBuying()
			bs := curr.BanknoteSelling()
			bb := curr.BanknoteBuying()

			if tc.code != curr.Code() {
				t.Errorf("expected: %s\ngot: %s", tc.code, curr.Code())
			}
			if tc.forexSelling != fs {
				t.Errorf("expected: %s\ngot: %s", tc.forexSelling, fs)
			}
			if tc.forexBuying != fb {
				t.Errorf("expected: %s\ngot: %s", tc.forexBuying, fb)
			}
			if tc.banknoteSelling != bs {
				t.Errorf("expected: %s\ngot: %s", tc.banknoteSelling, bs)
			}
			if tc.banknoteBuying != bb {
				t.Errorf("expected: %s\ngot: %s", tc.banknoteBuying, bb)
			}
		})
	}
}

func TestTodayDate(t *testing.T) {
	// TODO Hardcoded date will be problem.
	today := "18042023"
	got := todayDate()
	if today != got {
		t.Errorf("expected: %s\ngot: %s", today, got)
	}
	t.Logf("got: %s", got)
}

func TestNew(t *testing.T) {
	c := New()
	t.Log(c.date)
}
