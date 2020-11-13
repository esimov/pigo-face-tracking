package facetrack

import (
	"time"
)

var (
	pnx, pny, pt = 0, 0, time.Now()
)

// detectMovement capture the head movement and based on the velocity and direction
// it will trigger the corresponding keystroke event.
func (c *Canvas) detectMovement(nx, ny, th int) string {
	var cmd string
	ct := time.Now()
	dt := ct.Sub(pt).Seconds()

	dx, dy := nx-pnx, ny-pny
	//if time.Since(pt).Seconds()+0.4 > time.Since(ct).Seconds() {
	if (abs(dx) > th || abs(dy) > th) && dt < 1 {
		if abs(dx) > abs(dy) {
			if dx < 0 {
				cmd = "right"
			} else {
				cmd = "left"
			}
		} else {
			if dy > 0 {
				cmd = "down"
			} else {
				cmd = "up"
			}
		}
	}
	pnx, pny, pt = nx, ny, ct
	//}
	return cmd
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
