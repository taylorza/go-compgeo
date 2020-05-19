package quadtree

import "github.com/taylorza/go-compgeo/pkg/geom2d"

// QuadTree data structure used to index spacial data. The index opimizes queries looking for all the points that are within a specified area.
type QuadTree struct {
	root       *quadTreeNode
	maxPerNode int
	maxDepth   int
}

// Option is a function type that would be used to configure a quadtree.
type Option func(*QuadTree)

// MaxPerNode sets the maximum number of items in a quadtree node before the node is split.
func MaxPerNode(n int) Option {
	return func(q *QuadTree) {
		q.maxPerNode = n
	}
}

// MaxDepth sets the maximum depth after which a node will no longer be split.
func MaxDepth(n int) Option {
	return func(q *QuadTree) {
		q.maxDepth = n
	}
}

// New returns an instance of a QuadTree data structure that will index inserted items
// that have a location that falls within the passed rectangle. Options that can be set
// include the MaxDepth and MaxPerNode options which control the rate and maximum degree
// of subdivision of the nodes in the quadtree.
func New(rc geom2d.Rectangle, opts ...Option) *QuadTree {
	const (
		defaultMaxPerNode = 4
		defaultMaxDepth   = 4
	)

	q := &QuadTree{maxPerNode: defaultMaxPerNode, maxDepth: defaultMaxDepth}
	q.root = &quadTreeNode{q: q, rc: rc, depth: 0}

	for _, opt := range opts {
		opt(q)
	}

	return q
}

// Insert inserts and indexes a item satisfying the Locator interface into the quadtree.
func (q *QuadTree) Insert(p geom2d.Locator) bool {
	return q.root.insert(p)
}

// Query queries the quadtree for all items that fall within the boundary of the passed RangeMatcher.
func (q *QuadTree) Query(r geom2d.RangeMatcher) []geom2d.Locator {
	var result []geom2d.Locator
	q.root.query(r, &result)
	return result
}

// Rects returns all the rectangles of the quadtree
func (q *QuadTree) Rects() []geom2d.Rectangle {
	var queue []*quadTreeNode
	var rects []geom2d.Rectangle
	queue = append(queue, q.root)
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		rects = append(rects, n.rc)
		for _, c := range n.children {
			queue = append(queue, c)
		}
	}
	return rects
}

type quadTreeNode struct {
	q        *QuadTree
	rc       geom2d.Rectangle
	children []*quadTreeNode
	pts      []geom2d.Locator
	depth    int
}

func (n *quadTreeNode) insert(p geom2d.Locator) bool {
	if !n.rc.Contains(p) {
		return false
	}

	if (n.depth == n.q.maxDepth) || (len(n.pts) < n.q.maxPerNode && len(n.children) == 0) {
		n.pts = append(n.pts, p)
		return true
	}

	if len(n.children) == 0 {
		n.partition()
	}

	for i := 0; i < len(n.children); i++ {
		if n.children[i].insert(p) {
			return true
		}
	}

	return false
}

func (n *quadTreeNode) partition() {
	rects := partitionRect(n.rc)
	n.children = make([]*quadTreeNode, 4, 4)
	for i, rc := range rects {
		n.children[i] = &quadTreeNode{q: n.q, rc: rc, depth: n.depth + 1}
	}

	for _, pt := range n.pts {
		for _, c := range n.children {
			if c.insert(pt) {
				break
			}
		}
	}

	n.pts = nil
}

func partitionRect(r geom2d.Rectangle) []geom2d.Rectangle {
	rects := make([]geom2d.Rectangle, 4, 4)
	w2 := r.Width() / 2
	h2 := r.Height() / 2

	rects[0] = geom2d.NewRect(r.X(), r.Y(), w2, h2)
	rects[1] = geom2d.NewRect(r.X()+w2, r.Y(), w2, h2)
	rects[2] = geom2d.NewRect(r.X(), r.Y()+h2, w2, h2)
	rects[3] = geom2d.NewRect(r.X()+w2, r.Y()+h2, w2, h2)
	return rects
}

func (n *quadTreeNode) query(r geom2d.RangeMatcher, result *[]geom2d.Locator) {
	if !r.Intersects(n.rc) {
		return
	}

	for _, pt := range n.pts {
		if r.Contains(pt) {
			*result = append(*result, pt)
		}
	}

	for _, c := range n.children {
		c.query(r, result)
	}
}
