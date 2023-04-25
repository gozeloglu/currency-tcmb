package currency

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Currency stores its code, forex/banknote buying/selling information.
type Currency struct {
	code            Code
	unit            string
	forexBuying     string
	forexSelling    string
	banknoteBuying  string
	banknoteSelling string
}

// TCMB stores the forex/banknote buying/selling and date data for all currency.
type TCMB struct {
	currency map[Code]*Currency
	date     string
}

// tcmbXML is a type for parsing XML data.
type tcmbXML struct {
	XMLName      xml.Name   `xml:"Tarih_Date"`
	Text         string     `xml:",chardata"`
	DateTr       string     `xml:"Tarih,attr"`
	Date         string     `xml:"Date,attr"`
	BulletinNo   string     `xml:"Bulten_No,attr"`
	CurrencyList []currency `xml:"Currency"`
}

// currency is a type for parsing XML data.
type currency struct {
	Text            string `xml:",chardata"`
	CrossOrder      string `xml:"CrossOrder,attr"`
	Kod             string `xml:"Kod,attr"`
	CurrencyCode    string `xml:"CurrencyCode,attr"`
	Unit            string `xml:"Unit"`
	Name            string `xml:"Isim"`
	CurrencyName    string `xml:"CurrencyName"`
	ForexBuying     string `xml:"ForexBuying"`
	ForexSelling    string `xml:"ForexSelling"`
	BanknoteBuying  string `xml:"BanknoteBuying"`
	BanknoteSelling string `xml:"BanknoteSelling"`
	CrossRateUSD    string `xml:"CrossRateUSD"`
	CrossRateOther  string `xml:"CrossRateOther"`
}

const tcmbURL = "https://www.tcmb.gov.tr/kurlar"

var (
	ErrHTTP      = errors.New("failed to fetch data")
	ErrReadAll   = errors.New("failed to read all data")
	ErrUnmarshal = errors.New("failed to unmarshal XML data")
)

// New creates a currency object with all currencies prices. For now, it fetches
// today's prices.
func New(opt ...OptionFunc) *TCMB {
	t := &TCMB{currency: make(map[Code]*Currency)}

	for _, o := range opt {
		o(t)
	}

	// If custom date is not passed to parameter, set today's date.
	if t.date == "" {
		today := todayDate()
		t.date = today
	}
	term := termFrom(t.date)
	mbXML, err := fetchCurrency(term, t.date)
	if err != nil {
		return &TCMB{}
	}
	t.updateCurrencyMap(mbXML)
	return t
}

// FromCurrencyCode returns currency information data of the given currency code.
// For example, you can retrieve the currency information of the US Dollar with
// calling FromCurrencyCode(USD).
func (t *TCMB) FromCurrencyCode(code Code) *Currency {
	if len(t.currency) == 0 {
		return &Currency{}
	}
	return t.currency[code]
}

// ForexSelling returns forex selling for the Currency.
func (c *Currency) ForexSelling() string {
	return c.forexSelling
}

// ForexBuying returns forex buying for the Currency.
func (c *Currency) ForexBuying() string {
	return c.forexBuying
}

// BanknoteBuying returns banknote buying for the Currency.
func (c *Currency) BanknoteBuying() string {
	return c.banknoteBuying
}

// BanknoteSelling returns banknote selling for the Currency.
func (c *Currency) BanknoteSelling() string {
	return c.banknoteSelling
}

// Code returns Currency code.
func (c *Currency) Code() Code {
	return c.code
}

// fetchCurrency fetches the given term and date information.
func fetchCurrency(term string, date string) (tcmbXML, error) {
	u := fmt.Sprintf("%s/%s/%s.xml", tcmbURL, term, date)
	resp, err := http.Get(u)
	if err != nil {
		return tcmbXML{}, ErrHTTP
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return tcmbXML{}, ErrReadAll
	}
	var mbXML tcmbXML
	err = xml.Unmarshal(b, &mbXML)
	if err != nil {
		return tcmbXML{}, ErrUnmarshal
	}
	return mbXML, nil
}

// todayDate returns today's date in DDMMYYYY format.
func todayDate() string {
	year, month, day := time.Now().Date()
	m := fmt.Sprintf("%d", month)
	d := fmt.Sprintf("%d", day)
	if month < 10 {
		m = "0" + m
	}
	if day < 10 {
		d = "0" + d
	}
	return fmt.Sprintf("%d%s%d", day, m, year)
}

// termFrom converts date to YYYYMM format for the URL.
func termFrom(date string) string {
	month := date[2:4]
	year := date[4:]
	return fmt.Sprintf("%s%s", year, month)
}

// updateCurrencyMap sets Currency object to TCMB.currency map.
func (t *TCMB) updateCurrencyMap(mbXML tcmbXML) {
	for _, curr := range mbXML.CurrencyList {
		code := Code(curr.CurrencyCode)
		c := &Currency{
			code:            code,
			unit:            curr.Unit,
			forexBuying:     curr.ForexBuying,
			forexSelling:    curr.ForexSelling,
			banknoteBuying:  curr.BanknoteBuying,
			banknoteSelling: curr.BanknoteSelling,
		}
		t.currency[code] = c
	}
}
