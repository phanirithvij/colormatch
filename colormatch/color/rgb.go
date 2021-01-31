package color

import (
	"encoding/json"
	"fmt"
	"image/color"
	"math"

	"github.com/RobCherry/vibrant"
)

// RGBA non-pre alpha multiplied rgba color
type RGBA struct{ color.NRGBA }

// XYZ rgb to XYZ
func (rgb RGBA) XYZ() XYZ {
	newR := pivotForXYZ(rgb.R)
	newG := pivotForXYZ(rgb.G)
	newB := pivotForXYZ(rgb.B)
	return XYZ{
		X: (newR*0.4124 + newG*0.3576 + newB*0.1805) * 100,
		Y: (newR*0.2126 + newG*0.7152 + newB*0.0722) * 100,
		Z: (newR*0.0193 + newG*0.1192 + newB*0.9505) * 100,
	}
}

// Lab rgb to lab convertion
func (rgb RGBA) Lab() Lab {
	xyz := rgb.XYZ()
	newX := pivotForLab(xyz.X / 95.047)
	newY := pivotForLab(xyz.Y / 100.000)
	newZ := pivotForLab(xyz.Z / 108.883)

	l := math.Max(116*newY-16, 0)
	a := 500 * (newX - newY)
	b := 200 * (newY - newZ)

	return Lab{
		L: l,
		A: a,
		B: b,
	}
}

// HSB rgb to hsb convertion
func (rgb RGBA) HSB() HSB {
	rPrime := float64(rgb.R) / 255.0
	gPrime := float64(rgb.G) / 255.0
	bPrime := float64(rgb.B) / 255.0

	cMin := math.Min(math.Min(rPrime, gPrime), bPrime)
	cMax := math.Max(math.Max(rPrime, gPrime), bPrime)

	// Start with V. V is easy
	b := math.Round(cMax * 100)

	if cMax == 0 || cMax == cMin {
		return HSB{H: 0, S: 0, B: b}
	}

	// Next up, S
	delta := (cMax - cMin)
	s := math.Round(delta / cMax * 100)

	h := 0.0

	// Finally, H
	switch cMax {
	case rPrime:
		deltaInc := 0.0
		if gPrime < bPrime {
			deltaInc = 6
		}
		h = (gPrime-bPrime)/delta + deltaInc
	case gPrime:
		h = (bPrime-rPrime)/delta + 2
	case bPrime:
		h = (rPrime-gPrime)/delta + 4
	}

	// get h in degrees
	h = math.Round(h * 60)

	return HSB{H: h, S: s, B: b}
}

// Hex rgb to hex string
func (rgb RGBA) Hex() Hex {
	packed := rgb.Packed()
	// it is of the form 0xaarrggbb
	str := packed.String()
	rgbStr := str[4:]
	a := str[2:4]
	return Hex{h: "#" + rgbStr + a}
}

// RGB noop conversion
func (rgb RGBA) RGB() RGBA {
	return rgb
}

// RGBA image/color.Color compat
func (rgb RGBA) RGBA() (uint32, uint32, uint32, uint32) {
	return rgb.NRGBA.RGBA()
}

// String string representation
func (rgb RGBA) String() string {
	// https://stackoverflow.com/a/64306225/8608146
	bytes, err := json.Marshal(rgb)
	if err != nil {
		type temp RGBA
		return fmt.Sprintf("%+v\n", temp(rgb))
	}
	return string(bytes)
}

// Packed convert to vibrant rgbint struct
func (rgb RGBA) Packed() vibrant.RGBAInt {
	return vibrant.RGBAIntModel.Convert(rgb.NRGBA).(vibrant.RGBAInt)
}
