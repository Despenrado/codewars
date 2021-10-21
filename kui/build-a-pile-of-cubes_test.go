package kui

import "math"

func FindNb(m int) int {
	n := 1
	tempM := 1
	for m > tempM {
		n++
		tempM = tempM + int(math.Pow(float64(n), 3))
	}
	if tempM != m {
		return -1
	}
	return n
}
