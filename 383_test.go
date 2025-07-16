package main

import "testing"

func canConstruct(ransomNote string, magazine string) bool {
	var m [26]int
	for _, c := range magazine {
		m[c-'a']++
	}

	for _, c := range ransomNote {
		if m[c-'a'] <= 0 {
			return false
		}
		m[c-'a']--
	}
	return true
}

func TestCanConstruct(t *testing.T) {
	if canConstruct("a", "b") {
		t.Errorf("TEST 1 FAILD")
	}

	if canConstruct("aa", "ab") {
		t.Errorf("TEST 2 FAILED")
	}

	if !canConstruct("aa", "aab") {
		t.Errorf("TEST 3 FAILEDs")
	}
}
