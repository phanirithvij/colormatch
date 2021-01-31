package utils

// file: lib/find_furthest_colors.rb

// A quick method of getting the greatest possible distance between two colors,
// using Calculate::Distance. Important since Calculate::MatchScore
// is normalizing data using this range.

// It turns out the greatest distance, in LAB colorspace, is ~256,
// the distance between pure blue #0000FF and pure green #00FF00.

import "github.com/phanirithvij/colormatch/colormatch/color"

// MaxDist max distance data
type MaxDist struct {
	dist   float64
	colors []color.Color
}

func compareAllInDB() MaxDist {
	maxDist := 0.0
	maxColors := make([]color.Color, 2)

	colors := fetchColors()

	for _, c1 := range colors {
		for _, c2 := range colors {
			dist := Dist(c1, c2, "LAB")
			if dist > maxDist {
				maxDist = dist
				maxColors[0] = c1
				maxColors[1] = c2
			}
		}
	}

	return MaxDist{
		dist:   maxDist,
		colors: maxColors,
	}
}

func fetchColors() []color.Color {
	// TODO fetch from DB ?
	return make([]color.Color, 2)
}
