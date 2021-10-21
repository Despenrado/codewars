package kui

import (
	"fmt"
	"testing"
)

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}
	// v1*t + g = v2*t
	// g = v2*t - v1*t
	// g/t = v2 - v1
	// 1/t = (v2 - v1) / g
	// t = g / (v2 - v1)
	t := float64(float64(g) / float64(v2-v1))
	// return [3]int{int(t), int(t*60) % 60, int(math.Round(t*60*60)) % 60}
	return [3]int{int(t), int(t*60) % 60, int(t*60*60) % 60}
}

func TestRace(t *testing.T) {
	r := Race(80, 91, 37)
	fmt.Println(r)
}
