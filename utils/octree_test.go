package octree

import (
    "sol/oct"
    "testing"
    "github.com/go-utils/unum"
)

func TestOctreeInit (t *testing.T) {
    octree = &oct.Octree(unum.Vec3(0,0,0), unum.Vec(1,1,1))

}

func TestOctreeInsert (t *testing.T) {
    octree = &oct.Octree(unum.Vec3(0,0,0), unum.Vec(1,1,1))

}

func TestOctreeSearch (t *testing.T) {
    octree = &oct.Octree(unum.Vec3(0,0,0), unum.Vec(1,1,1))

}

func BenchmarkOctreeInsert (b *testing.B) {

}

func BenchmarkOctreeSearch (b *testing.B) {

}
