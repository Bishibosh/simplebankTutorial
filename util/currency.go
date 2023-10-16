package util

const (
	CAD = "CAD"
	EUR = "EUR"
	USD = "USD"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
