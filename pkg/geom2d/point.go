package geom2d

// Point represents a point with an X and Y coordinate. Point satisfies the Locator interface.
type Point struct {
	x, y float64
}

// NewPoint creates a Point with the specified X and Y coordinates
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// X returns the X coordinate of the point
func (p Point) X() float64 {
	return p.x
}

// Y returns the Y coordinate of the point
func (p Point) Y() float64 {
	return p.y
}
