package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	BRL = "BRL"
)

var currencies []string = []string{USD, EUR, CAD, BRL}

func IsSupportedCurrency(currency string) bool {
	for _, i := range currencies {
		if i == currency {
			return true
		}
	}

	return false
}
