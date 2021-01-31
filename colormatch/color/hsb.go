package color

import (
	"image/color"
	"math"
)

// file: app/services/colour/convert.rb

// HSB HSB  color format
type HSB struct {
	H, S, B float64
}

// RGB hsb to rgb color format
func (hsb HSB) RGB() RGBA {
	// Make sure our values are in the range 0-1 instead of 0-360 or 0-100
	h := hsb.H / 360.0
	s := hsb.S / 100.0
	v := hsb.B / 100.0
	// We're using V instead of B to avoid name conflicts with Blue in RGB.

	i := int(math.Floor(h * 6))
	f := (h * 6.0) - float64(i)
	p := v * (1 - s)
	q := v * (1 - f*s)
	t := v * (1 - (1-f)*s)

	r, g, b := .0, .0, .0

	switch i % 6 {
	case 0:
		r, g, b = v, t, p
	case 1:
		r, g, b = q, v, p
	case 2:
		r, g, b = p, v, t
	case 3:
		r, g, b = p, q, v
	case 4:
		r, g, b = t, p, v
	case 5:
		r, g, b = v, p, q
	}

	return RGBA{
		NRGBA: color.NRGBA{
			R: uint8(math.Floor(r * 255)),
			G: uint8(math.Floor(g * 255)),
			B: uint8(math.Floor(b * 255)),
			A: 0,
		},
	}
}

// Lab converts a color from hsb to RGB space
func (hsb HSB) Lab() Lab {
	// Let's first convert HSB to RGB, and then RGB to LAB
	return hsb.RGB().Lab()
}

// Hex convert to hex
func (hsb HSB) Hex() Hex {
	return hsb.RGB().Hex()
}

// RGBA image/color.Color compat
func (hsb HSB) RGBA() (uint32, uint32, uint32, uint32) {
	return hsb.RGB().RGBA()
}

// HSBModel is the color.Model for the HSB type.
var HSBModel = color.ModelFunc(func(c color.Color) color.Color {
	if _, ok := c.(HSB); ok {
		return c
	}
	nrgba := color.NRGBAModel.Convert(c).(color.NRGBA)
	return RGBA{NRGBA: nrgba}.HSB()
})

// HSB noop conversion
func (hsb HSB) HSB() HSB {
	return hsb
}
