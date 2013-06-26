
package primitives

import (
	"github.com/go-gl/gl"
)

func BuildPrimitive (prim func()) {
    // all primitives should be built at runtime (for a scene)
	object = gl.GenLists(1)
	gl.NewList(object, gl.COMPILE)
	gl.Materialfv(gl.FRONT, gl.AMBIENT_AND_DIFFUSE, []float32{0.8, 0.1, 0.0, 1.0})
	prim()
	gl.EndList()
    return object
}

func PlaceObject (compiledPrim gl.List, scale [3]double, translation [3]double, rotation [4]double) {
    // objects are placed when necessary
	gl.PushMatrix()
	gl.Scaled(scale)
	gl.Translated(translation)
	gl.Rotated(rotation)
	gl.CallList(compiledPrim)
	gl.PopMatrix()
}

func RelationalObject () {
    // object structure constructed of related items through a graph structure
    // scaling, translations, etc. can be computed by relative values
}

func Cube (x double, y double, z double, ) {
    //build with meshbuffer
	gl.Begin(gl.POLYGON)
    gl.Color3f(1.0, 1.0, 1.0)
    gl.Vertex3d(0.5, -0.5, 0.5)
    gl.Vertex3d(0.5, 0.5, 0.5)
    gl.Vertex3d(-0.5, 0.5, 0.5)
    gl.Vertex3d(-0.5, -0.5, 0.5)
    gl.End()
}

func Cylinder () {

}


