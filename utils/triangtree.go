package triangtree

import (
	"log"
	"fmt"
    "github.com/go-utils/unum"
)

type TriaPoint struct {
    Pos unum.Vec3
}

type TriaTree struct {
    origin unum.Vec3
    halfDimension unum.Vec3

    Children [8]TriaTree
    Point OctreePoint
}

// Accepts three points within the space
// Points form the initial layout of the dataset
func (oct *Octree) create () {

}

func (oct *Octree) isLeafNode () {
    return oct.children[0] == NULL
}

func (oct *Octree) containingPoint (point unum.Vec3) (int) {
    var oct int = 0
    switch {
    case point.X >= oct.origin.X: oct |= 4
    case point.Y >= oct.origin.Y: oct |= 2
    case point.Z >= oct.origin.Z: oct |= 1
    }
    return oct
}

func (oct *Octree) insert (point OctreeNode) {
    if(oct.isLeafNode()){
        if(oct.Point == NULL){
            oct.Point = point;
            return;
        } else {
            oldPoint := oct.Point
            var xSide float64
            var ySide float64
            var zSide float64

            for i := 0; i < 8 ; i++ {

                if (i&4) { xSide = .5 }else{ xSide = -.5 }
                if (i&2) { ySide = .5 }else{ ySide = -.5 }
                if (i&1) { zSide = .5 }else{ zSide = -.5 }

                newOrigin := &unum.Vec3{
                    X: halfDimension.X * xSide,
                    Y: halfDimension.Y * ySide,
                    Z: halfDimension.Z * zSide,
                }
                newOrigin.Add(oct.origin)
                oct.Children[i] = &Octree{
                    origin        : newOrigin,
                    halfDimension : halfDimension * float64(.5),
                }
            }
            oct.Children[oct.containingPoint(oldPoint.Pos)].insert(oldPoint)
            oct.Children[oct.containingPoint(point.Pos)].insert(point)
        }
    } else {
        oct.Children[oct.containingPoint(point.Pos)].insert(point)
    }
}

func (oct *Octree) getPointsBox (bmin unum.Vec3, bmax unum.Vec3, points []OctreePoint) {
    if(oct.isLeafNode()){
        if(oct.Point != NULL){
            p := oct.Point.Pos
            if(p.X > bmax.X || p.Y>bmax.Y || p.Z > bmax.Z) {return}
            if(p.X < bmin.X || p.Y<bmin.Y || p.Z < bmin.Z) {return}
            append(points, oct.Point)
        }
    } else {
        for i := 0 ; i<8 ; i++ {
            cmax := oct.Children[i].origin + oct.Children[i].halfDimension
            cmin := oct.Children[i].origin - oct.Children[i].halfDimension

            if(cmin.X > bmax.X || cmin.Y>bmax.Y || cmin.Z > bmax.Z) {continue}
            if(cmax.X < bmin.X || cmax.Y<bmin.Y || cmax.Z < bmin.Z) {continue}

            //race condition?
            go oct.Children[i].getPointsBox(bmin, bmax, points)
        }
    }
}
