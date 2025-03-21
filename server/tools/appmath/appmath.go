package appmath

import "math"

func RoundToDecimals(value float64, places int) float64 {
	factor := math.Pow(10, float64(places))
	return math.Round(value*factor) / factor
}

func CaclPercentFromSum(sum, percent float64) float64 {
	return sum * percent / 100
}
