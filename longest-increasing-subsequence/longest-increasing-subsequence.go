package longestincreasingsubsequence

func lengthOfLIS(num []int) int {
	switch len(num) {
	case 0:
		return 0
	case 1:
		return 1
	}

	length, first, _ := helper(num, 1)
	if num[0] < num[first] {
		return length + 1
	}
	return length
}

// 9
// 1, 3, 6, 7, 9, 4, 10, 5, 6
func helper(num []int, i int) (length, last_1, last int) {
	if i == len(num)-1 {
		return 1, i, i
	}
	le, la1, la2 := helper(num, i+1)
	if num[i] < num[la1] {
		return le + 1, i, la1
	} else if num[i] == num[la1] {
		return le, la1, la2
	} else if num[i] < num[la2] {
		return le, i, la2
	} else {
		return le, la1, la2
	}
}
