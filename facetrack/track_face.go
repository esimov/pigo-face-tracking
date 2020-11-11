package facetrack

import (
	"fmt"
	"time"
)

var (
	pnx, pny, pt = 0, 0, time.Now()
)

func (c *Canvas) detectMovement(nx, ny, th int) string {
	var pressedKey string
	ct := time.Now()
	dt := ct.Sub(pt).Seconds()

	dx, dy := nx-pnx, ny-pny
	if time.Since(pt).Seconds()+0.4 > time.Since(ct).Seconds() {
		if (abs(dx) > th || abs(dy) > th) && dt < 1 {
			fmt.Println(dx, dy)
			if dx > dy {
				if dx > 0 {
					pressedKey = "right"
				} else {
					pressedKey = "left"
				}
			} else {
				if dy > 0 {
					pressedKey = "down"
				} else {
					pressedKey = "up"
				}
			}
			pt = time.Now()
		}
		pnx, pny, pt = nx, ny, ct
	}
	//fmt.Println(dt)
	return pressedKey
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
