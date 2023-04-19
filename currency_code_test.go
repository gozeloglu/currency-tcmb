package currency

import "testing"

func TestCode_String(t *testing.T) {
	codes := []Code{USD, AUD, DKK, EUR, GBP, CHF, SEK, CAD, KWD, NOK, SAR, JPY, BGN, RON, RUB, IRR, CNY, PKR, QAR, KRW, AZN, AED, Code("ABC")}
	codesStr := []string{
		"US Dollar",
		"AUSTRALIAN DOLLAR",
		"DANISH KRONE",
		"EURO",
		"POUND STERLING",
		"SWISS FRANK",
		"SWEDISH KRONA",
		"CANADIAN DOLLAR",
		"KUWAITI DINAR",
		"NORWEGIAN KRONE",
		"SAUDI RIYAL",
		"JAPANESE YEN",
		"BULGARIAN LEV",
		"NEW LEU",
		"RUSSIAN ROUBLE",
		"IRANIAN RIAL",
		"CHINESE RENMINBI",
		"PAKISTANI RUPEE",
		"QATARI RIAL",
		"SOUTH KOREAN WON",
		"AZERBAIJANI NEW MANAT",
		"UNITED ARAB EMIRATES DIRHAM",
		"N/A"}

	for i, c := range codes {
		codeFull := c.String()
		if codeFull != codesStr[i] {
			t.Errorf("expected: %s\ngot: %s", codesStr[i], codeFull)
		}
	}
}
