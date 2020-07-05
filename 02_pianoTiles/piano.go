package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {

	x, y := robotgo.GetMousePos()
	println("Pos:", x, y)
	println("Color :", robotgo.GetPixelColor(x, y))

	xVal := []int{1371, 1521, 1629, 1764}

	for {
		for i := range xVal {
			if robotgo.GetPixelColor(xVal[i], 660) == "222222" {
				robotgo.Move(xVal[i], 660)
				robotgo.Click()
			}
		}
	}

}
