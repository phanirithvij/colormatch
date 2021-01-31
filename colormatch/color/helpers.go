package color

import (
	"math"
)

// file: app/services/colour/convert.rb

// pivotForLab helper method
func pivotForLab(n float64) float64 {
	if n > 0.008856 {
		return math.Pow(n, (1.0 / 3.0))
	}
	return ((903.3*n + 16) / 116)
}

// pivotForXYZ helper method
func pivotForXYZ(num uint8) float64 {
	n := float64(num)
	n /= 255.0
	if n > 0.04045 {
		return math.Pow(((n + 0.055) / 1.055), 2.4)
	}
	return n / 12.92
}
