package photo

import (
	"image/color"

	col "github.com/phanirithvij/colormatch/colormatch/color"
)

// file: app/services/photo/get_histogram_data.rb

// HistogramData histogram entry
type HistogramData struct {
	Count int
	RGB   color.NRGBA
	Hex   col.Hex
	HSB   col.HSB
	Lab   col.Lab
}
