package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	j := x % 10
	if j == 0 {
		return false
	}
	i := x / 10
	s := make([]uint8, 0)
	for i != 0 {
		x = i
		s = append(s, uint8(j))
		i = x / 10
		j = x % 10
	}
	s = append(s, uint8(j))
	i = 0
	j = len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
