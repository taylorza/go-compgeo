package quadtree

import (
	"math/rand"
	"testing"

	"github.com/taylorza/go-compgeo/pkg/geom2d"
)

func Test_QuadTreeCreateAndInsert(t *testing.T) {
	qt := New(geom2d.NewRect(0, 0, 100, 100), MaxPerNode(4), MaxDepth(4))

	if !qt.Insert(geom2d.NewPoint(5, 5)) {
		t.Fail()
	}

	if qt.Insert(geom2d.NewPoint(-5, -5)) {
		t.Fail()
	}
}

func Test_QuadTreeSubdivide(t *testing.T) {
	qt := New(geom2d.NewRect(0, 0, 100, 100), MaxPerNode(4), MaxDepth(4))

	qt.Insert(geom2d.NewPoint(25, 25))
	qt.Insert(geom2d.NewPoint(75, 25))
	qt.Insert(geom2d.NewPoint(25, 75))
	qt.Insert(geom2d.NewPoint(75, 75))
	qt.Insert(geom2d.NewPoint(55, 55))

	if len(qt.root.pts) != 0 || len(qt.root.children) == 0 {
		t.Fail()
	}

	if len(qt.root.children[0].pts) != 1 {
		t.Fail()
	}

	if len(qt.root.children[1].pts) != 1 {
		t.Fail()
	}

	if len(qt.root.children[2].pts) != 1 {
		t.Fail()
	}

	if len(qt.root.children[3].pts) != 2 {
		t.Fail()
	}
}

func Test_QuadTreeQuerySubdivided(t *testing.T) {
	qt := New(geom2d.NewRect(0, 0, 100, 100), MaxPerNode(4), MaxDepth(4))

	qt.Insert(geom2d.NewPoint(25, 25))
	qt.Insert(geom2d.NewPoint(75, 25))
	qt.Insert(geom2d.NewPoint(25, 75))
	qt.Insert(geom2d.NewPoint(75, 75))
	qt.Insert(geom2d.NewPoint(55, 55))

	r := qt.Query(geom2d.NewRect(0, 0, 100, 60))
	if len(r) != 3 {
		t.Fail()
	}

	r = qt.Query(geom2d.NewRect(50, 50, 50, 50))
	if len(r) != 2 {
		t.Fail()
	}
}

func Test_QuadTreeQueryCircle(t *testing.T) {
	qt := New(geom2d.NewRect(0, 0, 100, 100), MaxPerNode(4), MaxDepth(4))

	for x := 0; x < 100; x += 10 {
		for y := 0; y < 100; y += 10 {
			qt.Insert(geom2d.NewPoint(float32(x), float32(y)))
		}
	}

	r1 := qt.Query(geom2d.NewRect(0, 0, 100, 100))
	r2 := qt.Query(geom2d.NewCircle(50, 50, 50))

	if len(r1) == len(r2) {
		t.Fail()
	}

	if len(r2) == 0 {
		t.Fail()
	}
}

func Test_QuadTreeRects(t *testing.T) {
	qt := New(geom2d.NewRect(0, 0, 100, 100), MaxPerNode(4), MaxDepth(4))

	for x := 0; x < 100; x += 10 {
		for y := 0; y < 100; y += 10 {
			qt.Insert(geom2d.NewPoint(float32(x), float32(y)))
		}
	}

	r := qt.Rects()
	_ = r
}

func Benchmark_Quadtree(b *testing.B) {
	qt := New(geom2d.NewRect(0, 0, 100, 100))

	for i := 0; i < 1000; i++ {
		qt.Insert(geom2d.NewPoint(float32(rand.Intn(100)), float32(rand.Intn(100))))
	}
	for i := 0; i < b.N; i++ {
		for x := 0; x < 90; x += 10 {
			for y := 0; y < 90; y += 10 {
				qt.Query(geom2d.NewRect(float32(x), float32(y), 9, 9))
			}
		}
	}
}
