package ui

import (
    "container/list"
	"log"
	"fmt"
	"github.com/go-gl/glfw"
	"os"
)

type Window struct {
    width, height int
    title string
    MouseBtn func() (interface{})
    MouseWheel func() (interface{})
    MousePos func() (interface{})
    KeyEvt func() (interface{})
    CharEvt func() (interface{})
    // should we also have peek accessors?
    // these should be slices
    Render func()
    Start func()
}

type posEvent struct {
    x int
    y int
}

type btnEvent struct {
    button int
    state int
}

type deltaEvent struct {
    delta int
}

func CreateWindow(width int, height int, title string) (*Window){
	var err error
    window := &Window{width: width, height: height, title: title}
    // initialize logger
    logf, err := os.OpenFile("./logs/window.log", os.O_WRONLY|os.O_CREATE, 0640)
    logger := log.New(logf, "", log.Ldate|log.Ltime)

	if err = glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		return nil
	}

    // width, hieght, r,g,b,a (color depth bits), depth, stencil, mode
    if err = glfw.OpenWindow(window.width, window.height, 8, 8, 8, 0, 0, 0, glfw.Windowed); err != nil {
        fmt.Fprintf(os.Stderr, "[e] %v\n", err)
        return nil
    }

	// Enable vertical sync on cards that support it.
	glfw.SetSwapInterval(1)

	// Set window title
	glfw.SetWindowTitle(title)

    onResize := func(w, h int) {
        logger.Printf("resized: %dx%d\n", w, h)
    }

    onClose := func() int {
        logger.Println("window closed\n")
        return 1 // return 0 to keep window open.
    }

    // list callback generators
    createBtnList := func(btnNext func() (interface{}), device string) func(button, state int){
        btnList := list.New()
        btnNext = func() (interface{}) { return btnList.Remove(btnList.Back()) }
        return func(button, state int) {
            btnList.PushFront(btnEvent{button: button, state: state})
            logger.Printf("%s button: %d, %d\n", device, button, state)
        }
    }

    createPosList := func(posNext func() (interface{}), device string) func(x, y int){
        posList := list.New()
        posNext = func() (interface{}) { return posList.Remove(posList.Back()) }
        return func(x, y int) {
            posList.PushFront(posEvent{x: x, y: y})
            logger.Printf("%s pos: %d, %d\n", device, x, y)
        }
    }

    createDeltaList := func(deltaNext func() (interface{}), device string) func(delta int){
        deltaList := list.New()
        deltaNext = func() (interface{}) { return deltaList.Remove(deltaList.Back()) }
        return func(delta int) {
            deltaList.PushFront(deltaEvent{delta: delta})
            logger.Printf("%s delta: %d\n", device, delta)
        }
    }

	glfw.SetWindowSizeCallback(onResize)
	glfw.SetWindowCloseCallback(onClose)
	glfw.SetMousePosCallback(createPosList(window.MousePos, "mouse pos"))
	glfw.SetMouseButtonCallback(createBtnList(window.MouseBtn, "mouse"))
	glfw.SetMouseWheelCallback(createDeltaList(window.MouseWheel, "mouse wheel"))
	glfw.SetKeyCallback(createBtnList(window.KeyEvt, "keyboard"))
	glfw.SetCharCallback(createBtnList(window.CharEvt, "char"))

    window.Start = func () {
        defer glfw.Terminate()
        defer glfw.CloseWindow()
        running := true
        for running {
            window.Render()
            glfw.SwapBuffers()
            // Break out of loop when Escape key is pressed, or window is closed.
            running = glfw.Key(glfw.KeyEsc) == 0 && glfw.WindowParam(glfw.Opened) == 1
        }
    }

    return window
}
