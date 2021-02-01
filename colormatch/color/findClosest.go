package color

import (
	"math"

	"github.com/phanirithvij/colormatch/colormatch/models"
)

// FindClosest find the closest match
func FindClosest(color Color, useBins bool) Color {
	// We want to find which of the colors in our DB is closest to the provided color.
	// So, we need start by converting to LAB if it isn't already, for accuracy.
	// Then, do some pythagorean math.
	var bin models.Bin
	if useBins {
		if isGreyscale(color) {
			bin = models.Bin{}
			// bin =  Bin.first
		} else {
			bin = getNearestBin(color, nil)
		}
		return getNearestColor(color, bin.Colors)
	}
	return getNearestColor(color, []models.ColorModel{})
	// return getNearestColor(color, Color.all)
}

func isGreyscale(color Color) bool {
	return math.Round(color.Lab().A) == 0 && math.Round(color.Lab().B) == 0
}

func getNearestBin(color Color, bins []models.Bin) models.Bin {
	closest := math.Inf(+1)
	var nearestBin *models.Bin

	for _, b := range bins {
		dist := Dist(color, b.Exemplar, LabMode)
		if dist < closest {
			closest = dist
			nearestBin = &b
		}
	}

	return *nearestBin
}

func getNearestColor(c1 Color, binColors []models.ColorModel) Color {
	var closestColor *Color
	closestDist := math.Inf(+1)

	for _, c := range binColors {
		distance := Dist(c1, c, LabMode)
		if distance < closestDist {
			closestDist = distance
			x := Color(FromColorModel(c))
			closestColor = &x
		}
	}
	return *closestColor
}
