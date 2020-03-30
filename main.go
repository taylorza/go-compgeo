package main

import (
	"fmt"
	"math/rand"

	"github.com/taylorza/go-compgeo/pkg/prim2d"
	"github.com/taylorza/go-compgeo/pkg/quadtree"
)

func main() {

	// Create a new QuadTree that covers the area from (0, 0)-(100, 100)
	qt := quadtree.New(prim2d.NewRect(0, 0, 100, 100))

	// Add 1000 random items to the QuadTree
	for i := 0; i < 1000; i++ {
		qt.Insert(prim2d.NewPoint(float32(rand.Intn(100)), float32(rand.Intn(100))))
	}

	// Query the quadtree for all the items that are in the rectangle (10, 10) - (35, 25).
	// Note that the rectangle is created witha width and height of 25, 15, which give the bottom right coordinate of (35, 25)
	r1 := qt.Query(prim2d.NewRect(10, 10, 25, 15))

	// Query the quadtree for all the items that are within 20 units from the location 45, 45. This uses a circle to query the quadtree
	r2 := qt.Query(prim2d.NewCircle(45, 45, 20))

	fmt.Println(r1)
	fmt.Println(r2)
}
