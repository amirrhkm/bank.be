package util

const (
	USD = "USD"
	EUR = "EUR"
	MYR = "MYR"
	SGD = "SGD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, MYR, SGD:
		return true
	}
	return false
}
