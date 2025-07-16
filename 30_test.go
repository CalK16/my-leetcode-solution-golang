package main

import (
	"testing"
)

func all(s []bool) bool {
	for _, b := range s {
		if !b {
			return false
		}
	}
	return true
}

func permutation(words []string, selected []bool, prev string, res map[string]int) {
	if all(selected) {
		res[prev] = 1
		return
	}

	for i := range len(words) {
		if !selected[i] {
			selected[i] = true
			permutation(words, selected, prev+words[i], res)
			selected[i] = false
		}
	}
}

func findSubstring(s string, words []string) []int {
	n := len(words[0]) * len(words)
	m := make(map[string]int)
	permutation(words, make([]bool, len(words)), "", m)

	ans := make([]int, 0)
	for i := 0; i <= len(s)-n; i++ {
		_, ok := m[s[i:i+n]]
		if ok {
			ans = append(ans, i)
		}
	}

	return ans
}

func TestFindSubstring(t *testing.T) {
	ans := findSubstring("barfoothefoobarman", []string{"foo", "bar"})
	if !compareSlices(ans, []int{0, 9}) {
		t.Errorf("TEST 1 FAILED: %#v != %#v", ans, []int{0, 9})
	}

	ans = findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})
	if !compareSlices(ans, []int{}) {
		t.Errorf("TEST 2 FAILED")
	}

	ans = findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"})
	if !compareSlices(ans, []int{6, 9, 12}) {
		t.Errorf("TEST 3 FAILED")
	}

	ans = findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"})
	if !compareSlices(ans, []int{8}) {
		t.Errorf("TEST 4 FAILED")
	}
}
