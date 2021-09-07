package longestpalidrome

import "testing"

func TestMaxPalidrome(t *testing.T) {
	i, j := maxPalidrome("a", 0, 0)
	t.Logf("i: %d, j: %d", i, j)
	// t.Logf("sub: %s", "a"[i:j])
}

func TestLonestPalindrome(t *testing.T) {
	t.Log(longestPalidrom("babad"))
	t.Log(longestPalidrom("cbbd"))
	t.Log(longestPalidrom("a"))
	t.Log(longestPalidrom("ac"))
}
