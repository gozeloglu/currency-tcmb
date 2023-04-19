package currency

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Currency struct {
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
func New(currency string, opt ...OptionFunc) *Currency {
	// TODO Add functionality of fetching past date's prices.
	var c *Currency
	today := todayDate()
	term := termFrom(today)
	mbXML, err := fetchCurrency(term, today)
	if err != nil {
		return c
	}
	c.updateForexBuying(mbXML)
	c.updateForexSelling(mbXML)
	c.updateBanknoteBuying(mbXML)
	c.updateBanknoteSelling(mbXML)
	c.date = mbXML.Date
	return c
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
func (c *Currency) updateForexBuying(mbXML tcmbXML) {
	for _, curr := range mbXML.CurrencyList {
		c.forexBuying[Code(curr.CurrencyName)] = curr.ForexBuying
	}
}

// updateForexSelling sets forex selling information to Currency object.
func (c *Currency) updateForexSelling(mbXML tcmbXML) {
	for _, curr := range mbXML.CurrencyList {
		c.forexSelling[Code(curr.CurrencyName)] = curr.ForexSelling
	}
}

// updateBanknoteBuying sets banknote selling information to Currency object.
func (c *Currency) updateBanknoteBuying(mbXML tcmbXML) {
	for _, curr := range mbXML.CurrencyList {
		c.banknoteBuying[Code(curr.CurrencyName)] = curr.BanknoteBuying
	}
}

// updateBanknoteSelling sets banknote selling information to Currency object.
func (c *Currency) updateBanknoteSelling(mbXML tcmbXML) {
	for _, curr := range mbXML.CurrencyList {
		c.banknoteSelling[Code(curr.CurrencyName)] = curr.BanknoteSelling
	}
}
