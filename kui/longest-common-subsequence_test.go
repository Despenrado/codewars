package kui

import (
	"fmt"
	"regexp"
	"sync"
	"testing"
)

func LCS(s1, s2 string) string {
	longest := ""
	wg := sync.WaitGroup{}
	c := make(chan string)
	quit := make(chan int)
	go func() {
		tmp := ""
		for {
			select {
			case tmp = <-c:
				if len(tmp) > len(longest) {
					longest = tmp
				}
			case <-quit:
				return
			}
		}
	}()
	wg.Add(1)
	go genSubs(c, &wg, s1, 0, s2)
	wg.Wait()
	quit <- 0
	return longest
}

func genSubs(c chan string, wg *sync.WaitGroup, s string, n int, s2 string) {
	defer wg.Done()
	if n >= len(s) {
		// fmt.Println(s)
		if matched, _ := regexp.MatchString(makeRegexpString(s), s2); matched {
			c <- s
		}
		return
	}
	wg.Add(1)
	go genSubs(c, wg, s, n+1, s2)
	s = s[:n] + s[n+1:]
	wg.Add(1)
	go genSubs(c, wg, s, n, s2)
}

func makeRegexpString(s string) string {
	exp := ""
	for _, v := range s {
		exp += string(v) + ".*"
	}
	return exp
}

func TestLCS(t *testing.T) {
	fmt.Println(LCS("abcdef", "abc"))
}
