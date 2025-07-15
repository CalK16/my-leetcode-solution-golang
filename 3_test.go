package main

import (
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	var m [128]int
	ans := 0
	for r < len(s) {
		m[int(s[r])]++

		for m[int(s[r])] > 1 {
			m[int(s[l])]--
			l++
		}

		ans = max(ans, r-l+1)
		r++
	}

	return ans
}

func TestLengthOfLongestSubstring(t *testing.T) {
	if ans := lengthOfLongestSubstring("abcabcbb"); ans != 3 {
		t.Errorf("%d != 3", ans)
	}

	if ans := lengthOfLongestSubstring("bbbbb"); ans != 1 {
		t.Errorf("%d != 1", ans)
	}

	if ans := lengthOfLongestSubstring("pwwkew"); ans != 3 {
		t.Errorf("%d != 3", ans)
	}
}
