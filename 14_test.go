package main

import (
	"strings"
	"testing"
)

func longestCommonPrefix(strs []string) string {
	var builder strings.Builder
	for {
		idx := builder.Len()
		if len(strs[0]) <= idx {
			break
		}
		c := strs[0][idx]
		ok := true
		for _, s := range strs {
			if len(s) <= idx || s[idx] != c {
				ok = false
				break
			}
		}

		if !ok {
			break
		} else {
			builder.WriteByte(c)
		}
	}
	return builder.String()
}

func TestLonge4stCommonPrefix(t *testing.T) {
	if ans := longestCommonPrefix([]string{"flower", "flow", "flight"}); ans != "fl" {
		t.Errorf("%s != fl", ans)
	}

	if ans := longestCommonPrefix([]string{"dog", "racecar", "car"}); ans != "" {
		t.Errorf("%s != ''", ans)
	}

	if ans := longestCommonPrefix([]string{"ab", "a"}); ans != "a" {
		t.Errorf("%s != a", ans)
	}
}
