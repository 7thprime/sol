package bsp

import (
    "container/list"
	"log"
	"fmt"
	"github.com/go-gl/glfw"
	"os"
)

// classifyPoint accepts a Polygon to classify against and a Point to classify
func classifyPoint(poly Polygon, point Point) (string) {
    sideValue := poly.normal * point
    if(sideValue == poly.distance){
        return "coinciding"
    }else if(sideValue < poly.distance){
        return "behind"
    }else{
        return "infront"
    }
}



