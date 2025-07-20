package main

import (
	"strings"
	"testing"
)

func lengthOfLastWord(s string) int {
	tmp := strings.Trim(s, " ")
	ans := 0
	for i := len(tmp) - 1; i >= 0; i-- {
		if (65 <= tmp[i] && tmp[i] <= 90) || (97 <= tmp[i] && tmp[i] <= 122) {
			ans += 1
		} else {
			break
		}
	}
	return ans
}

func TestLengthOfLastWord(t *testing.T) {
	if ans := lengthOfLastWord("Hello World"); ans != 5 {
		t.Errorf("%d != 5", ans)
	}

	if ans := lengthOfLastWord("   fly me   to   the moon  "); ans != 4 {
		t.Errorf("%d != 4", ans)
	}

	if ans := lengthOfLastWord("luffy is still joyboy"); ans != 6 {
		t.Errorf("%d != 6", ans)
	}
}
