package renderer

import (
    "math"
    "github.com/go-utils/unum"
)

type Object3d struct {
    Pos       unum.Vec3
    dirty     bool
}

type DirectionalObject3d struct {
    Pos       unum.Vec3
    rollAngle float64
    yawAngle  float64
    UpAxis    unum.Vec3 //the vector considered Up
    dirty     bool
}

// Point DirectionalObject3d at a point
func (diObj *DirectionalObject3d) LookAt (point unum.Vec3) {

}

