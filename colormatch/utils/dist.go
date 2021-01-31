package utils

import (
	"math"

	"github.com/phanirithvij/colormatch/colormatch/color"
)

// Dist color distance
func Dist(c1, c2 color.Color, mode string) float64 {
	cx1, cx2 := c1, c2
	if mode == "Lab" {
		return labMath(cx1.Lab(), cx2.Lab())
	}
	return hsbMath(cx1.HSB(), cx2.HSB())
}

// hsbMath hsb distance
func hsbMath(c1, c2 color.HSB) float64 {
	return math.Sqrt(math.Pow(c1.H-c2.H, 2) + math.Pow(c1.S-c2.S, 2) + math.Pow(c1.B-c2.B, 2))
}

// labMath lab distance
func labMath(c1, c2 color.Lab) float64 {
	// L is weighted less 0.7
	return math.Sqrt(math.Pow(c1.L-c2.L, 2)*0.7 + math.Pow(c1.A-c2.A, 2) + math.Pow(c1.B-c2.B, 2))
}
