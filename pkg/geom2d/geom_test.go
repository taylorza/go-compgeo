package geom2d

import (
	"testing"
)

func Test_PointConstructor(t *testing.T) {
	pt := NewPoint(24, 42)
	if pt.X() != 24 || pt.Y() != 42 {
		t.Fail()
	}
}

func Test_CircleConstructor(t *testing.T) {
	c := NewCircle(0, 0, 50)
	if c.X() != 0 || c.Y() != 0 || c.R() != 50 {
		t.Fail()
	}
}

func Test_CircleContains(t *testing.T) {
	c := NewCircle(0, 0, 50)

	if !c.Contains(NewPoint(30, 30)) {
		t.Fail()
	}

	if !c.Contains(NewPoint(-30, -30)) {
		t.Fail()
	}

	if c.Contains(NewPoint(51, 51)) {
		t.Fail()
	}

	if c.Contains(NewPoint(-51, 50)) {
		t.Fail()
	}
}

func TestCircleRectIntersect(t *testing.T) {
	c := NewCircle(0, 0, 50)

	if !c.Intersects(NewRect(-5, -5, 10, 10)) {
		t.Fail()
	}

	if !c.Intersects(NewRect(35, -45, 100, 100)) {
		t.Fail()
	}

	if !c.Intersects(NewRect(-35, 45, 100, 100)) {
		t.Fail()
	}

	if c.Intersects(NewRect(51, 51, 10, 10)) {
		t.Fail()
	}
}

func Test_RectangleConstructor(t *testing.T) {
	c := NewRect(0, 0, 50, 60)
	if c.X() != 0 || c.Y() != 0 || c.Width() != 50 || c.Height() != 60 {
		t.Fail()
	}
}

func Test_RectangleLeftRightTopBottom(t *testing.T) {
	c := NewRect(50, 50, -10, 10)
	if c.Left() != 40 || c.Top() != 50 || c.Right() != 50 || c.Bottom() != 60 {
		t.Fail()
	}
}

func Test_RectangleContains(t *testing.T) {
	c := NewRect(0, 0, 50, 60)

	if !c.Contains(NewPoint(30, 30)) {
		t.Fail()
	}

	if c.Contains(NewPoint(-30, -30)) {
		t.Fail()
	}

	if c.Contains(NewPoint(51, 51)) {
		t.Fail()
	}

	if c.Contains(NewPoint(58, 48)) {
		t.Fail()
	}
}

func Test_RectangleRectangleIntersect(t *testing.T) {
	c := NewRect(0, 0, 50, 60)

	if !c.Intersects(NewRect(-5, -5, 10, 10)) {
		t.Fail()
	}

	if !c.Intersects(NewRect(35, -45, 100, 100)) {
		t.Fail()
	}

	if !c.Intersects(NewRect(-35, 45, 100, 100)) {
		t.Fail()
	}

	if c.Intersects(NewRect(51, 51, 10, 10)) {
		t.Fail()
	}
}
