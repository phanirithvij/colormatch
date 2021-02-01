package photo

// Thresholds
const (
	BrightnessThreshold = 18
	SaturationThreshold = 18
)

// ValidHue if a valid hue
// Because not all colours have an accurate hue (for example, according to hue, white is red).
// This service returns a boolean based on whether the hue is 'valid', and should be considered,
// based on the saturation and brightness of the colour.
func ValidHue(color HistEntry) bool {
	return color.HSB.B > BrightnessThreshold && color.HSB.S > SaturationThreshold
}
