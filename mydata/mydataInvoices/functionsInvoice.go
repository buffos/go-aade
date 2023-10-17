package mydataInvoices

import (
	"math"
)

func roundToMoney(v float64) float64 {
	return math.Round(v*100) / 100
}
