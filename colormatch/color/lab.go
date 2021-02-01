package color

// Lab lab color format
type Lab struct {
	L, A, B float64
}

// Lab noop conversion
func (l Lab) Lab() Lab {
	return l
}

// TODO lab to rgba
