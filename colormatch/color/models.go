package color

import "image/color"

// Model can convert any Color to one from its own color model. The conversion
// may be lossy.
type Model interface {
	Convert(c Color) Color
}

// ModelFunc returns a Model that invokes f to implement the conversion.
func ModelFunc(f func(Color) Color) Model {
	// Note: using *modelFunc as the implementation
	// means that callers can still use comparisons
	// like m == RGBAModel. This is not possible if
	// we use the func value directly, because funcs
	// are no longer comparable.
	return &modelFunc{f}
}

type modelFunc struct {
	f func(Color) Color
}

func (m *modelFunc) Convert(c Color) Color {
	return m.f(c)
}

// LocalRGBAModel is the color.Model for the RGBA type.
var LocalRGBAModel = ModelFunc(func(c Color) Color {
	if _, ok := c.(RGBA); ok {
		return c
	}
	nrgba := color.NRGBAModel.Convert(c).(color.NRGBA)
	return RGBA{NRGBA: nrgba}
})

// LocalHSBModel is the color.Model for the HSB type.
var LocalHSBModel = ModelFunc(func(c Color) Color {
	if _, ok := c.(HSB); ok {
		return c
	}
	rgba := LocalRGBAModel.Convert(c).(RGBA)
	return rgba.HSB()
})

// LocalHexModel is the color.Model for the Hex type.
var LocalHexModel = ModelFunc(func(c Color) Color {
	if _, ok := c.(HSB); ok {
		return c
	}
	rgba := LocalRGBAModel.Convert(c).(RGBA)
	return rgba.Hex()
})
