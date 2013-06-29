package model_obj

import (
    "bufio"
    "io"
    "os"
    "strings"
	"github.com/go-gl/gl"
)


func LoadFileTypeObj(filename string) {
    rdr := bufio.NewReader(os.Stdin)
    switch line, err := rdr.ReadString('\n'); err {
    case nil:
        switch line {
        case strings.HasPrefix(line, "v "): //vertex
            verticies := strings.Split(line[2:], " ")
            gl.Vertex4f(float64(verticies[0]), float64(verticies[1]), float64(verticies[2]), float64(1.0))
        case strings.HasPrefix(line, "f "): //face
            verticies := strings.Split(line[2:], " ")
            gl.GLushort 
            a--; b--; c--;

        case strings.HasPrefix(line, "#"):
        }
    case io.EOF:
    default:
    };


}

