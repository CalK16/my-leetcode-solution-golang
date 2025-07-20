package main

import (
	"strings"
	"testing"
)

func reverseWords(s string) string {
	tmp := strings.Trim(s, " ")
	lst := strings.Split(tmp, " ")
	var builder strings.Builder
	for i := len(lst) - 1; i >= 0; i-- {
		word := lst[i]
		if word == "" {
			continue
		}
		builder.Write([]byte(word))
		if i != 0 {
			builder.WriteByte(' ')
		}
	}
	return builder.String()
}

func TestReverseWords(t *testing.T) {
	if ans := reverseWords("the sky is blue"); ans != "blue is sky the" {
		t.Errorf("%s != \"blue is sky the\"", ans)
	}

	if ans := reverseWords("  hello world  "); ans != "world hello" {
		t.Errorf("%s != \"world hello\"", ans)
	}

	if ans := reverseWords("a good   example"); ans != "example good a" {
		t.Errorf("%s != \"example good a\"", ans)
	}
}
