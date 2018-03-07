package utilities

import "math"

func Round(value float64, places int ) float64 {
	var round float64
	roundOn := 0.5
	pow := math.Pow(10, float64(places))
	digit := pow * value
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}