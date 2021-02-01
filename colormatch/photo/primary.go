package photo

import "github.com/phanirithvij/colormatch/colormatch/color"

// PrimaryColor struct
type PrimaryColor struct {
	Type  string
	Color color.Color
	Count int
}

// ExtractPrimaryColour extract the primary color from the histogram data
func ExtractPrimaryColour(colors []HistEntry) PrimaryColor {
	return matchColorTodb(colors[0])
}
func matchColorTodb(c HistEntry) PrimaryColor {
	return PrimaryColor{
		Type:  "primary",
		Color: color.FindClosest(c.RGB, false),
		Count: c.Count,
	}
}
