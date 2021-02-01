package photo

import (
	"math"

	"github.com/phanirithvij/colormatch/colormatch/maths"
)

// Stats stats
type Stats struct {
	HSB hsbStats
	Lab labStats
}

type hsbStats struct {
	H, S, B chStats
}

type labStats struct {
	L, A, B chStats
}

type chStats struct {
	Mean      float64
	Deviation float64
}

// Channel enum
type Channel uint8

const (
	h Channel = iota
	s
	v
)

const (
	l Channel = iota
	a
	b
)

// Space colorspace enum
type Space uint8

const (
	hsb Space = iota
	lab
)

// GetStats get the stats
func GetStats(colors []HistEntry) Stats {
	return Stats{
		HSB: hsbStats{
			H: getChannelStats(colors, hsb, h),
			S: getChannelStats(colors, hsb, s),
			B: getChannelStats(colors, hsb, v),
		},
		Lab: labStats{
			L: getChannelStats(colors, lab, l),
			A: getChannelStats(colors, lab, a),
			B: getChannelStats(colors, lab, b),
		},
	}
}

func getChannelStats(colors []HistEntry, space Space, channel Channel) chStats {
	data := buildRepresentativeArray(colors, space, channel)
	mean, deviation := 0.0, 0.0
	// greyscale image
	if channel == h && len(colors) == 0 {
		mean, deviation = 0, 0
	} else {
		mean = maths.Mean(data)
		deviation = maths.StandardDeviation(data)
	}
	return chStats{
		Mean:      mean,
		Deviation: deviation,
	}
}

// I want to take into account how many times a colour is repeated, when getting deviation.
// However, it adds too much to the processing time to use every pixel, so I'm dividing
// occurances by 500 so get a general representation of the occurances, without killing CPU.
func buildRepresentativeArray(colors []HistEntry, space Space, channel Channel) []float64 {
	// White and black both have a hue of 0, which is in the reds.
	// A blue image with black, therefore, would have a purple mean. This is bad.
	// We want to ignore images whose brightness is either very very low or very very high.
	if channel == h {
		colors = removeGreyAndDimHues(colors)
	}

	data := make([]float64, len(colors))
	for i, c := range colors {
		// TODO use reflect instead of this massive switch case (?)
		// https://stackoverflow.com/a/18931036/8608146
		switch space {
		case hsb:
			switch channel {
			case h:
				data[i] = c.HSB.H * math.Ceil(float64(c.Count)/500.0)
			case s:
				data[i] = c.HSB.S * math.Ceil(float64(c.Count)/500.0)
			case v:
				data[i] = c.HSB.B * math.Ceil(float64(c.Count)/500.0)
			}
		case lab:
			switch channel {
			case l:
				data[i] = c.Lab.L * math.Ceil(float64(c.Count)/500.0)
			case a:
				data[i] = c.Lab.A * math.Ceil(float64(c.Count)/500.0)
			case b:
				data[i] = c.Lab.B * math.Ceil(float64(c.Count)/500.0)
			}
		}
	}

	return data
}

func removeGreyAndDimHues(colors []HistEntry) []HistEntry {
	filtered := []HistEntry{}
	for _, c := range colors {
		if ValidHue(c) {
			filtered = append(filtered, c)
		}
	}
	return filtered
}
