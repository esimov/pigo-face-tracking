package pigo

import "math"

// abs returns the absolute value of the provided number
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// min returns the minum value between two numbers
func min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}

// max returns the maximum value between two numbers
func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

// round returns the nearest integer, rounding ties away from zero.
func round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}
