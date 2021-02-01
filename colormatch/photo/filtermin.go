package photo

const (
	// DefaultMin default min
	DefaultMin = 0.001
)

// FilterColorsByOccurance filter colors by occurance
func FilterColorsByOccurance(data HistData, minimum float64) []HistEntry {
	if minimum <= 0 {
		minimum = DefaultMin
	}
	minOccurance := float64(data.Pixels) * minimum

	colors := []HistEntry{}
	for _, c := range data.Colors {
		if float64(c.Count) >= minOccurance {
			colors = append(colors, c)
		}
	}
	return colors
}
