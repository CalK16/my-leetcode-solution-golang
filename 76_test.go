package main

import (
	"testing"
)

// contains() return true if a contains all keys that b has, and the value of the key is larger
// than ones that are in b
func contains(a, b map[byte]int) bool {
	if len(a) < len(b) {
		return false
	}
	for key, bval := range b {
		aval, ok := a[key]
		if !ok || bval > aval {
			return false
		}
	}

	return true
}

// makeMap() turns `t` to a map with key being its characters and value being the count of characters.
func makeMap(t string) map[byte]int {
	count := make(map[byte]int)
	for _, c := range t {
		_, ok := count[byte(c)]
		if !ok {
			count[byte(c)] = 0
		}
		count[byte(c)]++
	}
	return count
}

func minWindow(s string, t string) string {
	n, m := len(s), len(t)
	if n < m {
		return ""
	}

	ans := ""
	minLen := n + 1
	count := makeMap(t)

	curr := make(map[byte]int)
	l, r := 0, 0
	for r < n {
		_, ok := curr[s[r]]
		if !ok {
			curr[s[r]] = 0
		}

		curr[s[r]]++

		for contains(curr, count) {
			if minLen > r-l+1 {
				ans = s[l : r+1]
				minLen = len(ans)
			}
			curr[s[l]]--
			l++
		}

		r++
	}

	if minLen == n+1 {
		return ""
	}
	return ans
}

func TestMinWindow(t *testing.T) {
	if ans := minWindow("ADOBECODEBANC", "ABC"); ans != "BANC" {
		t.Errorf("TEST 1 FAILED: %s != %s", ans, "BANC")
	}

	if ans := minWindow("a", "a"); ans != "a" {
		t.Errorf("TEST 2 FAILED")
	}

	if ans := minWindow("a", "aa"); ans != "" {
		t.Errorf("TEST 3 FAILED")
	}
}
