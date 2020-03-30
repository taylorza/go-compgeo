package prim2d

// Rectangle represents a rectangle. Rectangle satisfies the RangeMatcher interface.
type Rectangle struct {
	x, y, w, h float32
}

// NewRect creates a new rectangle with the specified X, Y coordinates and width and height
func NewRect(x, y, w, h float32) Rectangle {
	return Rectangle{x: x, y: y, w: w, h: h}
}

// X returns the X coordinate of the rectangle
func (r Rectangle) X() float32 {
	return r.x
}

// Y returns the Y coordinate of the rectangle
func (r Rectangle) Y() float32 {
	return r.y
}

// Width returns the width of the rectangle
func (r Rectangle) Width() float32 {
	return r.w
}

// Height returns the height of the rectangle
func (r Rectangle) Height() float32 {
	return r.h
}

// Left returns the left most X coordinate of the rectangle
func (r Rectangle) Left() float32 {
	return min(r.x, r.x+r.w)
}

// Right returns the right most X coordinate of the rectangle
func (r Rectangle) Right() float32 {
	return max(r.x, r.x+r.w)
}

// Top returns the top most Y coordinate of the rectangle
func (r Rectangle) Top() float32 {
	return min(r.y, r.y+r.h)
}

// Bottom returns the bottom most Y coordinate of the rectangle
func (r Rectangle) Bottom() float32 {
	return max(r.y, r.y+r.h)
}

// Contains returns true if the passed locator is within the boundaries of the rectangle.
func (r Rectangle) Contains(p Locator) bool {
	return r.Left() <= p.X() && r.Right() > p.X() &&
		r.Top() <= p.Y() && r.Bottom() > p.Y()
}

// Intersects returns true if the two rectangles intersect each other
func (r Rectangle) Intersects(o Rectangle) bool {
	return r.Left() < o.Right() && r.Right() > o.Left() &&
		r.Top() < o.Bottom() && r.Bottom() > o.Top()
}
