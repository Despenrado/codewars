package kui

import "testing"

func DigitalRoot(n int) int {
	if n < 10 {
		return n
	}
	res := 0
	for n > 0 {
		res += n % 10
		n = n / 10
	}
	return DigitalRoot(res)
}

func TestDigitalRoot(t *testing.T) {
	if DigitalRoot(942) != 6 {
		t.Failed()
	}
}
