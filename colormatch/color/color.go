package color

// Color color allows cross conversion
type Color interface {
	Hex() Hex
	HSB() HSB
	RGB() RGBA
	Lab() Lab
	// makes it compatible with image/color.Color interface
	RGBA() (r, g, b, a uint32)
}
