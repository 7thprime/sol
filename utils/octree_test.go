package octree

import (
    "sol/oct"
    "rand"
    "testing"
    "github.com/go-utils/unum"
)

// Helper: -1,1 random float64
func rand11(){
    return float64(-1.0) + (2 * rand.Float64())
}

// Helper: random unum.Vec3 
func newRandomPoint () {
    return &unum.Vec3{
        X: rand11(),
        Y: rand11(),
        Z: rand11(),
    }
}

// Globally scoped octree
var octree Octree

// Testing Octree initialization
func TestOctreeInit (t *testing.T) {
    octree = &oct.Octree(unum.Vec3(0,0,0), unum.Vec(1,1,1))
}

// Testing Octree insertion of new points
// TODO: test total number of points to ensure they were all inserted
func TestOctreeInsert (t *testing.T) {
    points := make([]unum.Vec3, 3)
    numRandomPoints = 1 * 1000 * 1000
    for i :=  0 ; i<numRandomPoints ; i++ {
        append(points, newRandomPoint())
    }
    for i :=  0 ; i<numRandomPoints ; i++ {
        octree.Insert(&OctreePoint{Pos: points[i]})
    }
}

// Testing Octree search within bounds
// TODO: deterministic test of number of points found within box
func TestOctreeSearch (t *testing.T) {
    qmin := &unum.Vec3{
        X: -.05,
        Y: -.05,
        Z: -.05,
    }
    qmax := &unum.Vec3{
        X: .05,
        Y: .05,
        Z: .05,
    }

    results := make([]OctreePoint, 3)
    octree.GetPointsInBox(qmin, qmax, results)
}

func BenchmarkOctreeInsert (b *testing.B) {

}

func BenchmarkOctreeSearch (b *testing.B) {

}

func BenchmarkNaiveSearch (b *testing.B) {

}
