package kui

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func PrimeFactors(n int) string {
	p := make(map[int]int)
	primeFactor(n, p)
	return parser(p)
}

func primeFactor(n int, p map[int]int) {
	if n == 1 {
		return
	}
	i := n - 1
	for n%i != 0 {
		i--
	}
	if i == 1 {
		p[n]++
		return
	}
	primeFactor(i, p)
	primeFactor(n/i, p)
}

func parser(p map[int]int) string {
	s := ""
	keys := []int{}
	for key := range p {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, v := range keys {
		if p[v] != 1 {
			s += "(" + strconv.Itoa(v) + "**" + strconv.Itoa(p[v]) + ")"
		} else {
			s += "(" + strconv.Itoa(v) + ")"
		}
	}
	return s
}
func TestPrimeFactors(t *testing.T) {
	fmt.Println(PrimeFactors(86240))
}
