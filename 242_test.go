package main

import "testing"

func isAnagram(s string, t string) bool {
	curr := make(map[rune]int)
	for _, c := range s {
		curr[c]++
	}

	for _, c := range t {
		curr[c]--
	}

	for _, val := range curr {
		if val != 0 {
			return false
		}
	}
	return true
}

func TestIsAnagram(t *testing.T) {
	if !isAnagram("anagram", "nagaram") {
		t.Errorf("TEST 1 FAILED")
	}

	if isAnagram("rat", "car") {
		t.Errorf("TEST 2 FAILED")
	}
}
