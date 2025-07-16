package main

import "testing"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s2t := make(map[rune]byte)
	t2s := make(map[rune]byte)

	for i, c := range s {
		_, found := s2t[c]
		if !found {
			if _, found2 := t2s[rune(t[i])]; found2 {
				return false
			}
			s2t[c] = t[i]
			t2s[rune(t[i])] = byte(c)
		} else if s2t[c] != t[i] || t2s[rune(t[i])] != byte(c) {
			return false
		}
	}
	return true
}

func TestIsIsmorphic(t *testing.T) {
	if !isIsomorphic("egg", "add") {
		t.Errorf("TEST 1 FAILED")
	}

	if isIsomorphic("foo", "bar") {
		t.Errorf("TEST 2 FAILED")
	}

	if isIsomorphic("badc", "baba") {
		t.Errorf("TEST 3 FAILED")
	}
}
