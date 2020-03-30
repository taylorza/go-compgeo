package prim2d

// Circle represents a circle. Circle satifies the RangeMatcher interface.
type Circle struct {
	x, y, r float32
}

// NewCircle creates a circle object.
func NewCircle(x, y, r float32) Circle {
	return Circle{x: x, y: y, r: r}
}

// Contains returns true if the passed locator is within the boundaries of the circle.
func (c Circle) Contains(p Locator) bool {
	dx := c.x - p.X()
	dy := c.y - p.Y()
	return (c.r * c.r) >= (dx*dx)+(dy*dy)
}

// Intersects returns true if the circle intersects with the passed rectangle.
func (c Circle) Intersects(o Rectangle) bool {
	closestX := clamp(c.x, o.x, o.x+o.w)
	closestY := clamp(c.y, o.y, o.y+o.h)
	dx := c.x - closestX
	dy := c.y - closestY
	d2 := (dx * dx) + (dy * dy)
	return d2 < (c.r * c.r)
}

// X returns the X coordinate of the circle's center point.
func (c Circle) X() float32 {
	return c.x
}

// Y returns the Y coordinate of the circle's center point.
func (c Circle) Y() float32 {
	return c.y
}

// R returns the radius of the circle.
func (c Circle) R() float32 {
	return c.r
}
