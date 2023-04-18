package currency

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Currency struct {
	currency        string
	forexBuying     map[string]float64
	forexSelling    map[string]float64
	banknoteBuying  map[string]float64
	banknoteSelling map[string]float64
	date            string
}

type TCMBXML struct {
	XMLName    xml.Name `xml:"Tarih_Date"`
	Text       string   `xml:",chardata"`
	DateTr     string   `xml:"Tarih,attr"`
	Date       string   `xml:"Date,attr"`
	BulletinNo string   `xml:"Bulten_No,attr"`
	Currency   []struct {
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
	} `xml:"Currency"`
}

const tcmbURL = "https://www.tcmb.gov.tr/kurlar"

// New creates a currency object with all currencies prices. For now, it fetches
// today's prices.
func New(currency string, opt ...OptionFunc) Currency {
	// TODO Add functionality of fetching past date's prices.
	today := todayDate()
	term := termFrom(today)
	tcmbXML, err := fetchCurrency(term, today)
	if err != nil {
		return Currency{}
	}
	fmt.Println(tcmbXML.Currency[0].CurrencyName)
	fmt.Println(tcmbXML.Currency[0].ForexSelling)
	return Currency{
		currency:        currency,
		forexBuying:     map[string]float64{},
		forexSelling:    map[string]float64{},
		banknoteBuying:  map[string]float64{},
		banknoteSelling: map[string]float64{},
		date:            "",
	}
}

// fetchCurrency fetches the given term and date information.
func fetchCurrency(term string, date string) (TCMBXML, error) {
	u := fmt.Sprintf("%s/%s/%s.xml", tcmbURL, term, date)
	resp, err := http.Get(u)
	if err != nil {
		// TODO Custom error can be returned
		return TCMBXML{}, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		// TODO Custom error can be returned
		return TCMBXML{}, err
	}
	var tcmbXML TCMBXML
	err = xml.Unmarshal(b, &tcmbXML)
	if err != nil {
		// TODO Custom error can be returned
		return TCMBXML{}, err
	}
	return tcmbXML, nil
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
