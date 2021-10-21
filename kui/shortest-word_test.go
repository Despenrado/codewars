package kui

import (
	"fmt"
	"strings"
	"testing"
)

func FindShort(s string) int {
	strs := strings.Split(s, " ")
	n := len(strs[0])
	for _, v := range strs {
		tmp := len(v)
		if n > tmp {
			n = tmp
		}
	}
	return n
}

func TestFindShort(t *testing.T) {
	r := FindShort("lol wtf it is complitely \n shit task")
	fmt.Println(r)
}
