package kui

import (
	"fmt"
	"math"
	"testing"
)

func KPrimesStep(k int, step int, m int, n int) [][]int {
	kPrimes := make([][]int, 0)

	for i := m; i+step <= n; i++ {
		if i != 0 && len(findPrimes(i, make([]int, 0), k)) == k && len(findPrimes(i+step, make([]int, 0), k)) == k {
			kPrimes = append(kPrimes, []int{i, i + step})
		}
	}
	if len(kPrimes) == 0 {
		return nil
	}
	return kPrimes
}

func findPrimes(number int, primes []int, maxPrimes int) []int {
	if number == 1 || len(primes) > maxPrimes {
		return primes
	}
	i := int(math.Sqrt(float64(number)))
	for number%i != 0 {
		i--
	}
	if i == 1 {
		return append(primes, number)
	}
	return findPrimes(number/i, findPrimes(i, primes, maxPrimes), maxPrimes)
}

func TestKPrimesStep(t *testing.T) {
	primes := KPrimesStep(3, 6, 26956, 27220)
	fmt.Println(primes)
	fmt.Println(findPrimes(27016, make([]int, 0), 3+1))

}
