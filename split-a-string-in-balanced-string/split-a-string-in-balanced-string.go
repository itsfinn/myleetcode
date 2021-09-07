package splitastringinbalancedstring

func balancedStringSplit(s string) int {
	var g, cnt int

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'R':
			g++
		case 'L':
			g--
		}
		if g == 0 {
			cnt++
		}
	}
	return cnt
}
