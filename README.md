# currency-tcmb

`currency-tcmb` is a Go API for the Central Bank of the Republic of Türkiye(TCMB) currency data. It fetches the data from the [URL](https://www.tcmb.gov.tr/kurlar/today.xml), 
parses the XML, and returns necessary information. It returns today's data. 

## Installation

```shell
go get github.com/gozeloglu/currency-tcmb
```

## Usage

```go
func main() {
    tcmb := currency.New()  // It fetches and parses the XML data in background
    usdCurrency := tcmb.FromCurrencyCode(currency.USD)  // Get USD currency against TRY  
    fmt.Println(usdCurrency.BanknoteBuying())
    fmt.Println(usdCurrency.BanknoteSelling())
}
```

You can retrieve the historical currency by passing `WithDate()` option. 

```go
tcmb := currency.New(WithDate(1, time.September, 2021))  // It fetches 01 September 2021 currency data.
```

## LICENSE
[MIT](LICENSE)