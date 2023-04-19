package currency

type Code string

const (
	USD = Code("USD")
	AUD = Code("AUD")
	DKK = Code("DKK")
	EUR = Code("EUR")
	GBP = Code("GBP")
	CHF = Code("CHF")
	SEK = Code("SEK")
	CAD = Code("CAD")
	KWD = Code("KWD")
	NOK = Code("NOK")
	SAR = Code("SAR")
	JPY = Code("JPY")
	BGN = Code("BGN")
	RON = Code("RON")
	RUB = Code("RUB")
	IRR = Code("IRR")
	CNY = Code("CNY")
	PKR = Code("PKR")
	QAR = Code("QAR")
	KRW = Code("KRW")
	AZD = Code("AZD")
	AED = Code("AED")
)

// String returns currency codes' full name.
func (c Code) String() string {
	switch c {
	case "USD":
		return "US Dollar"
	case "AUD":
		return "AUSTRALIAN DOLLAR"
	case "DKK":
		return "DANISH KRONE"
	case "EUR":
		return "EURO"
	case "GBP":
		return "POUND STERLING"
	case "CHF":
		return "SWISS FRANK"
	case "SEK":
		return "SWEDISH KRONA"
	case "CAD":
		return "CANADIAN DOLLAR"
	case "KWD":
		return "KUWAITI DINAR"
	case "NOK":
		return "NORWEGIAN KRONE"
	case "SAR":
		return "SAUDI RIYAL"
	case "JPY":
		return "JAPANESE YEN"
	case "BGN":
		return "BULGARIAN LEV"
	case "RON":
		return "NEW LEU"
	case "RUB":
		return "RUSSIAN ROUBLE"
	case "IRR":
		return "IRANIAN RIAL"
	case "CNY":
		return "CHINESE RENMINBI"
	case "PKR":
		return "PAKISTANI RUPEE"
	case "QAR":
		return "QATARI RIAL"
	case "KRW":
		return "SOUTH KOREAN WON"
	case "AZN":
		return "AZERBAIJANI NEW MANAT"
	case "AED":
		return "UNITED ARAB EMIRATES DIRHAM"
	default:
		return "N/A"
	}
}
