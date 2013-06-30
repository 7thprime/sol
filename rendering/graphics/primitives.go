package graphics

import (

)

func Cube (x double)  {
    //build with meshbuffer
	gl.Begin(gl.POLYGON)
    gl.Color3f(1.0, 1.0, 1.0)
    gl.Vertex3d(0.5, -0.5, 0.5)
    gl.Vertex3d(0.5, 0.5, 0.5)
    gl.Vertex3d(-0.5, 0.5, 0.5)
    gl.Vertex3d(-0.5, -0.5, 0.5)
    gl.End()
}

func Sphere (x) {

}

func Cylinder (x) {

}
