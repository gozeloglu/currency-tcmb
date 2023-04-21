package currency

import (
	"testing"
)

func TestTCMB_FromCurrencyCode(t *testing.T) {
	testCases := []struct {
		name            string
		code            Code
		expCode         Code
		forexBuying     string
		forexSelling    string
		banknoteBuying  string
		banknoteSelling string
		tcmbIsEmpty     bool
	}{
		{
			name:        "Empty TCMB.currency list",
			code:        USD,
			tcmbIsEmpty: true,
		},
		{
			name:            "USD",
			code:            USD,
			expCode:         USD,
			forexBuying:     "10",
			forexSelling:    "11",
			banknoteBuying:  "12",
			banknoteSelling: "13",
		},
		{
			name:            "EUR",
			code:            EUR,
			expCode:         EUR,
			forexBuying:     "15",
			forexSelling:    "16",
			banknoteBuying:  "17",
			banknoteSelling: "18",
		},
	}

	for _, tc := range testCases {
		tcmb := &TCMB{
			currency: map[Code]*Currency{
				USD: {
					code:            "USD",
					unit:            "1",
					forexBuying:     "10",
					forexSelling:    "11",
					banknoteBuying:  "12",
					banknoteSelling: "13",
				},
				EUR: {
					code:            "EUR",
					unit:            "1",
					forexBuying:     "15",
					forexSelling:    "16",
					banknoteBuying:  "17",
					banknoteSelling: "18",
				},
			},
			date: "20231904",
		}
		t.Run(tc.name, func(t *testing.T) {
			if tc.tcmbIsEmpty {
				tcmb = &TCMB{currency: nil, date: ""}
			}
			curr := tcmb.FromCurrencyCode(tc.code)
			fs := curr.ForexSelling()
			fb := curr.ForexBuying()
			bs := curr.BanknoteSelling()
			bb := curr.BanknoteBuying()

			if tc.expCode != curr.Code() {
				t.Errorf("expected: %s\ngot: %s", tc.expCode, curr.Code())
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

func TestFetchCurrency(t *testing.T) {
	testCases := []struct {
		name            string
		term            string
		date            string
		forexBuying     string
		forexSelling    string
		banknoteSelling string
		banknoteBuying  string
		expErr          bool
		err             error
	}{
		{
			name:            "Fetch past time currency rate",
			term:            "202304",
			date:            "04042023",
			forexBuying:     "19.2001",
			forexSelling:    "19.2347",
			banknoteBuying:  "19.1867",
			banknoteSelling: "19.2636",
		},
		{
			name:   "Fetch past time currency rate in weekend",
			term:   "202304",
			date:   "01042023",
			expErr: true,
			err:    ErrUnmarshal,
		},
		{
			name:   "Future time currency rate",
			term:   "203307",
			date:   "01072023",
			expErr: true,
			err:    ErrUnmarshal,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mbXML, err := fetchCurrency(tc.term, tc.date)
			if tc.expErr && err != tc.err {
				t.Fatalf("expected error %v\nactual error: %v", tc.err, err)
			}
			if err != nil {
				// if error is returned to avoid fail in following if checks
				return
			}
			if mbXML.CurrencyList[0].CurrencyName != "US DOLLAR" {
				t.Errorf("expected name: %s\nactual name: %s", "US DOLLAR", mbXML.CurrencyList[0].CurrencyName)
			}
			if mbXML.CurrencyList[0].ForexBuying != tc.forexBuying {
				t.Errorf("expected forex buying: %s\nactual forex buying: %s", tc.forexBuying, mbXML.CurrencyList[0].ForexBuying)
			}
			if mbXML.CurrencyList[0].ForexSelling != tc.forexSelling {
				t.Errorf("expected forex selling: %s\nactual forex selling: %s", tc.forexSelling, mbXML.CurrencyList[0].ForexSelling)
			}
			if mbXML.CurrencyList[0].BanknoteBuying != tc.banknoteBuying {
				t.Errorf("expected banknote buying: %s\nactual banknote buying: %s", tc.banknoteBuying, mbXML.CurrencyList[0].BanknoteBuying)
			}
			if mbXML.CurrencyList[0].BanknoteSelling != tc.banknoteSelling {
				t.Errorf("expected banknote selling: %s\nactual banknote selling: %s", tc.banknoteSelling, mbXML.CurrencyList[0].BanknoteSelling)
			}
		})
	}
}

func TestCurrency_ForexBuying(t *testing.T) {
	c := &Currency{
		code:        "EUR",
		forexBuying: "15.324",
	}

	fb := c.ForexBuying()
	if fb != "15.324" {
		t.Errorf("expected: 15.324\ngot: %s", fb)
	}
}

func TestCurrency_ForexSelling(t *testing.T) {
	c := &Currency{
		code:         "EUR",
		forexSelling: "14.987",
	}

	fs := c.ForexSelling()
	if fs != "14.987" {
		t.Errorf("expected: 14.987\ngot: %s", fs)
	}
}

func TestCurrency_BanknoteBuying(t *testing.T) {
	c := &Currency{
		code:           "EUR",
		banknoteBuying: "15.487",
	}

	bb := c.BanknoteBuying()
	if bb != "15.487" {
		t.Errorf("expected: 15.487\ngot: %s", bb)
	}
}

func TestCurrency_BanknoteSelling(t *testing.T) {
	c := &Currency{
		code:            "EUR",
		banknoteSelling: "15.426",
	}

	bs := c.BanknoteSelling()
	if bs != "15.426" {
		t.Errorf("expected: 15.426\ngot: %s", bs)
	}
}

func TestTermFrom(t *testing.T) {
	date := "04062023"
	term := termFrom(date)

	if term != "202306" {
		t.Errorf("expected: 202306\ngot: %s", term)
	}
}
