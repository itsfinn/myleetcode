package maximumproductsubarray

import "github.com/itsfinn/myleetcode/util"

func MaxProductSubstring(a []float64) float64 {
	maxEnd := a[0]
	minEnd := a[0]
	maxResult := a[0]
	length := len(a)
	for i := 1; i < length; i++ {
		end1 := maxEnd * a[i]
		end2 := minEnd * a[i]
		maxEnd = util.MaxFloat64(util.MaxFloat64(end1, end2), a[i])
		minEnd = util.MinFloat64(util.MinFloat64(end1, end2), a[i])
		maxResult = util.MaxFloat64(maxResult, maxEnd)
	}
	return maxResult
}
