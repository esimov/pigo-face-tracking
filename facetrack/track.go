package facetrack

import (
	"time"
)

const idleTime = 0.35

var (
	pnx, pny, pt, it = 0, 0, time.Now(), time.Now()
)

// detectMovement: captures the head movement and based on its
// velocity and direction it will trigger the corresponding keystroke event.
func (c *Canvas) detectMovement(nx, ny, th int) string {
	var cmd string
	ct := time.Now()
	dt := ct.Sub(pt).Seconds()

	dx, dy := nx-pnx, ny-pny
	if (abs(dx) > th || abs(dy) > th) && dt < 1 {
		// Trigger only one keystroke event in a certain time interval in case
		// the velocity of the face movement is greather than the predefined thershold value.
		if time.Since(it).Seconds() > idleTime {
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
			it = time.Now()
		}
	}
	pnx, pny, pt = nx, ny, ct

	return cmd
}

// abs returns the absolute value of the variable.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
