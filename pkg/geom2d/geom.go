package geom2d

// Locator interface represents any object that can be located by an X and Y coordinate
type Locator interface {
	//
	X() float64
	Y() float64
}

// RangeMatcher interface represents objects that define ranges that will be used to query spacial data.
type RangeMatcher interface {
	Contains(p Locator) bool
	Intersects(r Rectangle) bool
}

// Utility functions used in the primitives
func clamp(v, min, max float64) float64 {
	if v < min {
		v = min
	} else if v > max {
		v = max
	}
	return v
}

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}
