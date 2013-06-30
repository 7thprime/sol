package renderer

import (
    "bufio"
    "container/list"
	"github.com/go-gl/gl"
)


// Renderer 
type Renderer struct {
    queue *list.List
}

// Init initializes a renderer r
func (r *Renderer) Init() *Renderer {
    r.queue = list.New()
    return r
}

// New returns an initialized renderer.
func New () *Renderer { return new(Renderer).Init() }

// Render function for graphics loop
func (renderer *Renderer) Render () {
    renderer.queue.Front()
}
