package maths

import "math"

const (
	// RadMultiplier radian multiplier
	RadMultiplier = 57.29577951580927
	// DegMultiplier degree multiplier
	DegMultiplier = 0.0174532925
)

// PolarCoord r, theta
type PolarCoord struct {
	R     float64
	Angle float64
}

// CartesianCoord x, y
type CartesianCoord struct {
	X float64
	Y float64
}

// Mean mean of an array
func Mean(a []float64) float64 {
	sum := 0.0
	for _, v := range a {
		sum += v
	}
	return sum / float64(len(a))
}

// SampleVariance variance
func SampleVariance(a []float64) float64 {
	m := Mean(a)
	sum := 0.0
	for _, v := range a {
		sum += math.Pow(v-m, 2)
	}
	return sum / float64(len(a)-1)
}

// StandardDeviation standard deviation
func StandardDeviation(a []float64) float64 {
	return math.Sqrt(SampleVariance(a))
}

// ZScore z score
func ZScore(a float64, sample []float64) float64 {
	// raise "Need either the sample, or the mean and deviation, to calculate z-score" unless sample || (mean && deviation)
	// Calculate mean and deviation unless provided
	mean := Mean(sample)
	deviation := StandardDeviation(sample)
	return (a - mean) / deviation
}

// CircularMean circular mean for polarcoords
func CircularMean(a []float64, cartesianArr []CartesianCoord) float64 {
	if cartesianArr == nil {
		cartesianArr = make([]CartesianCoord, len(a))
		// a is an array of degree values from 0 to 360.
		// I need to convert it from degrees to cartesian coordinates [x, y].
		for i, angle := range a {
			cartesianArr[i] = ConvertToCartesian(1, angle)
		}
	}
	// cartesianArr is a 2-dimensional array with x-y values.
	// eg. [ [1, 4], [2, 3], [1, 0.5] ]
	// Let's get the average x and y
	xs := make([]float64, len(cartesianArr))
	ys := make([]float64, len(cartesianArr))
	for i := range cartesianArr {
		xs[i] = cartesianArr[i].X
		ys[i] = cartesianArr[i].Y
	}
	averagex := Mean(xs)
	averagey := Mean(ys)

	// let's convert this x,y pair back to polar coordinates, for our solution.
	polar := ConvertToPolar(averagex, averagey)
	return polar.Angle
}

// CircularDeviation circular deviation of polarcoords
func CircularDeviation(a []float64) float64 {
	cartesianArr := make([]CartesianCoord, len(a))
	for i, angle := range a {
		cartesianArr[i] = ConvertToCartesian(1, angle)
	}

	xs := make([]float64, len(cartesianArr))
	ys := make([]float64, len(cartesianArr))
	for i := range cartesianArr {
		xs[i] = cartesianArr[i].X
		ys[i] = cartesianArr[i].Y
	}
	// So this is tricky. We have two values for every point, [x, y].
	// Should I get the standard deviation of both x and y, and average them out?
	devx := StandardDeviation(xs)
	devy := StandardDeviation(ys)

	polar := ConvertToPolar(devx, devy)
	return polar.Angle
}

// ConvertToCartesian convert polar to cartesian coordinates
func ConvertToCartesian(radius float64, angle float64) CartesianCoord {
	// If supplied with the angle in degrees, we need to convert to radians.
	angle *= DegMultiplier

	// We're using the unit circle by funcault for all calculations; radius = 1
	return CartesianCoord{radius * math.Cos(angle), radius * math.Cos(angle)}
}

// ConvertToPolar convert cartesian to polar coordinates
func ConvertToPolar(x, y float64) PolarCoord {
	// we don't need the radius for calculating circular_mean, but what the hell.
	r := math.Sqrt(x*x + y*y)

	radians := math.Atan(y / x)
	angle := radians * RadMultiplier

	switch FindQuadrant(x, y) {
	case 2:
	case 3:
		angle += 180
	case 4:
		angle += 360
	}
	return PolarCoord{
		R:     r,
		Angle: angle,
	}
}

// FindQuadrant find quadrant of x,y
func FindQuadrant(x, y float64) int {
	if x > 0 {
		if y > 0 {
			return 1
		}
		return 4
	}
	if y > 0 {
		return 2
	}
	return 3
}
