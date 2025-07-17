package main

import (
	"testing"
)

func findSubstring(s string, words []string) []int {
	ans := make([]int, 0)
	n, m := len(words), len(words[0])
	wc := make(map[string]int)

	for _, word := range words {
		wc[word]++
	}

	for i := range len(s) - n*m + 1 {
		w := s[i : i+(n*m)]

		curr := make(map[string]int)
		for j := 0; j < len(w); j += m {
			curr[w[j:j+m]]++
		}

		same := true
		for word := range curr {
			val, ok := wc[word]
			if !ok || curr[word] != val {
				same = false
				break
			}
		}

		if same {
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
