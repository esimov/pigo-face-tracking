package facetrack

import (
	"time"
)

const idleTime = 0.3

var (
	pnx, pny, pt, lt = 0, 0, time.Now(), time.Now()
)

// detectMovement capture the head movement and based on the velocity and direction
// it will trigger the corresponding keystroke event.
func (c *Canvas) detectMovement(nx, ny, th int) string {
	var cmd string
	ct := time.Now()
	dt := ct.Sub(pt).Seconds()

	dx, dy := nx-pnx, ny-pny
	if (abs(dx) > th || abs(dy) > th) && dt < 1 {
		if (time.Since(lt).Seconds()) > idleTime {
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
			lt = time.Now()
		}
	}
	pnx, pny, pt = nx, ny, ct

	return cmd
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
