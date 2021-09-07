package longestpalidrome

import "fmt"

func longestPalidrom(s string) string {
	var maxi, maxj, i1, j1, i2, j2 int
	if len(s) <= 1 {
		return s
	}
	for i := 0; i < len(s)-1; i++ {
		i1, j1 = maxPalidrome(s, i, i)
		i2, j2 = maxPalidrome(s, i, i+1)
		if j2-i2 > j1-i1 && j2-i2 > maxj-maxi {

			maxi, maxj = i2, j2
		} else if j1-i1 > maxj-maxi {

			maxi, maxj = i1, j1
		}
	}
	fmt.Printf("i: %d, j: %d\n", maxi, maxj)
	return s[maxi:maxj]
}

func maxPalidrome(s string, i, j int) (int, int) {

	for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {

	}

	if i == j {
		return i, j + 1
	} else if i == j-1 {
		return i, j
	} else {
		return i + 1, j

	}
}
