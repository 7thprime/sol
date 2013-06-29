// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows how to set up a minimal GLFW application.
package main

import (
    "sol/ui"
)

func main() {
    win := ui.CreateWindow(600, 400, "test")
    win.Render(func (){

    })
    win.Start()
}
