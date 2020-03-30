package prim2d

import "math"

// Locator interface represents any object that can be located by an X and Y coordinate
type Locator interface {
	//
	X() float32
	Y() float32
}

// RangeMatcher interface represents objects that define ranges that will be used to query spacial data.
type RangeMatcher interface {
	Contains(p Locator) bool
	Intersects(r Rectangle) bool
}

func clamp(v, min, max float32) float32 {
	if v < min {
		v = min
	} else if v > max {
		v = max
	}
	return v
}

func min(x, y float32) float32 {
	if x < y {
		return x
	}
	return y
}

func max(x, y float32) float32 {
	if x > y {
		return x
	}
	return y
}

func eq(x, y float32) bool {
	return math.Abs(float64(x-y)) > math.SmallestNonzeroFloat64
}