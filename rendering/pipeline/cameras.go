package renderer

import (

)

type Perspective struct{
    Enabled bool
    FovY struct {
        Deg float64
        RadHalf float64
    }
    ZFar float64
    ZNear float64
}

// stores current perspective position of the scene
type Camera struct {
    Perspective renderer.Perspective
    Object3D Object3D
}
