package models

import (
	"math"

	"github.com/phanirithvij/colormatch/colormatch/color"
)

// Bin bin
type Bin struct {
	Exemplar ColorModel
	Colors   []ColorModel
}

// FindClosest find closest bin
func FindClosest(c ColorModel, bins []Bin) Bin {
	if bins == nil {
		// Let's grab all our bins (with eager-loaded exemplars) if not supplied
		// bins = Bin.includes(:exemplar).all
	}

	return getNearestBin(c, bins)
}

func getNearestBin(c ColorModel, bins []Bin) Bin {
	closest := math.Inf(+1)
	var nearestBin *Bin

	for _, b := range bins {
		dist := color.Dist(color.FromColorModel(c), color.FromColorModel(b.Exemplar), color.LabMode)
		if dist < closest {
			closest = dist
			nearestBin = &b
		}
	}

	return *nearestBin
}
