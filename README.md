# Computational Geometry Library 
[![Build Status](https://travis-ci.org/taylorza/go-compgeo.svg?branch=master)](https://travis-ci.org/taylorza/go-compgeo) [![Coverage Status](https://coveralls.io/repos/github/taylorza/go-compgeo/badge.svg?branch=master)](https://coveralls.io/github/taylorza/go-compgeo?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/taylorza/go-compgeo)](https://goreportcard.com/report/github.com/taylorza/go-compgeo) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/taylorza/go-compgeo)

Library with a few useful computational geometry routines

* QuadTree - 2d Spacial Index that can be used to quickly query for points within a area. This is great for optimizing collision detection between lots of objects, using the quadtree and help you quickly find objects close to you and only test those for collisions rather than test every object against every other object. For moving objects the quadtree is efficient enough that it can be rebuilt for each update frame.

## Examples

### Create a QuadTree and query it using rectangular and circular query boundaries
```go
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
```

## Copyright 
Copyright (C)2013-2020 by Chris Taylor (taylorza)
See [Licence](https://github.com/taylorza/go-compgeo/blob/master/LICENSE)
