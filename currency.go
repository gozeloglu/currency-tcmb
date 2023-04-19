package currency

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Currency stores its code, forex/banknote buying/selling information.
type Currency struct {
	code            Code
	forexBuying     string
	forexSelling    string
	banknoteBuying  string
	banknoteSelling string
}

// TCMB stores the forex/banknote buying/selling and date data for all currency.
type TCMB struct {
	forexBuying     map[Code]string
	forexSelling    map[Code]string
	banknoteBuying  map[Code]string
	banknoteSelling map[Code]string
	date            string
}

type tcmbXML struct {
	XMLName      xml.Name   `xml:"Tarih_Date"`
	Text         string     `xml:",chardata"`
	DateTr       string     `xml:"Tarih,attr"`
	Date         string     `xml:"Date,attr"`
	BulletinNo   string     `xml:"Bulten_No,attr"`
	CurrencyList []currency `xml:"Currency"`
}

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

// New creates a currency object with all currencies prices. For now, it fetches
// today's prices.
func New(opt ...OptionFunc) *TCMB {
	// TODO Add functionality of fetching past date's prices.
	var t *TCMB
	today := todayDate()
	term := termFrom(today)
	mbXML, err := fetchCurrency(term, today)
	if err != nil {
		return t
	}
	t.updateForexBuying(mbXML)
	t.updateForexSelling(mbXML)
	t.updateBanknoteBuying(mbXML)
	t.updateBanknoteSelling(mbXML)
	t.date = mbXML.Date
	return t
}

func (t *TCMB) FromCurrencyCode(code Code) *Currency {
	return &Currency{
		code:            code,
		forexBuying:     t.forexBuying[code],
		forexSelling:    t.forexSelling[code],
		banknoteBuying:  t.banknoteBuying[code],
		banknoteSelling: t.banknoteSelling[code],
	}
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
		// TODO Custom error can be returned
		return tcmbXML{}, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		// TODO Custom error can be returned
		return tcmbXML{}, err
	}
	var mbXML tcmbXML
	err = xml.Unmarshal(b, &mbXML)
	if err != nil {
		// TODO Custom error can be returned
		return tcmbXML{}, err
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

// updateForexBuying sets forex buying information to Currency object.
func (t *TCMB) updateForexBuying(mbXML tcmbXML) {
	for _, curr := range mbXML.CurrencyList {
		t.forexBuying[Code(curr.CurrencyName)] = curr.ForexBuying
	}
}

// updateForexSelling sets forex selling information to Currency object.
func (t *TCMB) updateForexSelling(mbXML tcmbXML) {
	for _, curr := range mbXML.CurrencyList {
		t.forexSelling[Code(curr.CurrencyName)] = curr.ForexSelling
	}
}

// updateBanknoteBuying sets banknote selling information to Currency object.
func (t *TCMB) updateBanknoteBuying(mbXML tcmbXML) {
	for _, curr := range mbXML.CurrencyList {
		t.banknoteBuying[Code(curr.CurrencyName)] = curr.BanknoteBuying
	}
}

// updateBanknoteSelling sets banknote selling information to Currency object.
func (t *TCMB) updateBanknoteSelling(mbXML tcmbXML) {
	for _, curr := range mbXML.CurrencyList {
		t.banknoteSelling[Code(curr.CurrencyName)] = curr.BanknoteSelling
	}
}
