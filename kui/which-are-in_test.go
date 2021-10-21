package kui

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	rm := make(map[string]byte)
	for _, v1 := range array2 {
		for _, v2 := range array1 {
			if strings.Contains(v1, v2) {
				rm[v2] = 1
			}
		}
	}
	r := []string{}
	for key := range rm {
		r = append(r, key)
	}
	sort.Strings(r)
	return r
}
