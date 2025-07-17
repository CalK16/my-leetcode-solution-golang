package main

import (
	"strings"
	"testing"
)

func wordPattern(pattern string, s string) bool {
	sl := strings.Split(s, " ")

	if len(sl) != len(pattern) {
		return false
	}

	p2s := make(map[byte]string)
	s2p := make(map[string]byte)
	for i := range pattern {
		word, p2sOK := p2s[pattern[i]]
		patt, s2pOK := s2p[sl[i]]
		if p2sOK && s2pOK {
			if s2p[word] != patt || p2s[patt] != word {
				return false
			}
		} else if !p2sOK && !s2pOK {
			s2p[sl[i]] = pattern[i]
			p2s[pattern[i]] = sl[i]
		} else {
			return false
		}

	}
	return true
}

func TestWordPattern(t *testing.T) {
	if !wordPattern("abba", "dog cat cat dog") {
		t.Errorf("TEST 1 FAILED")
	}

	if wordPattern("abba", "dog cat cat fish") {
		t.Errorf("TEST 2 FAILED")
	}

	if wordPattern("aaaa", "dog cat cat dog") {
		t.Errorf("TEST 3 FAILED")
	}
}
