// +build js,wasm

package main

import (
	"fmt"

	"github.com/esimov/pigo-face-tracking/tracking"
)

func main() {
	c := tracking.NewCanvas()
	webcam, err := c.StartWebcam()
	if err != nil {
		c.Alert("Webcam not detected!")
	} else {
		err := webcam.Render()
		if err != nil {
			c.Alert(fmt.Sprint(err))
		}
	}
}
