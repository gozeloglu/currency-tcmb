# currency-tcmb

`currency-tcmb` is a Go API for Turkish Central Bank currency data. It fetches the data from the [URL](https://www.tcmb.gov.tr/kurlar/today.xml), 
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

## LICENSE
[MIT](LICENSE)