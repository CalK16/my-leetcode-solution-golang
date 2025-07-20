package main

import (
	"strings"
	"testing"
)

func convert(s string, numRows int) string {
	builders := make([]strings.Builder, numRows)

	i, flag := 0, -1
	for _, c := range s {
		builders[i].WriteRune(c)
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}

	var ans strings.Builder
	for _, builder := range builders {
		ans.WriteString(builder.String())
	}
	return ans.String()
}

func TestConvert(t *testing.T) {
	if ans := convert("PAYPALISHIRING", 3); ans != "PAHNAPLSIIGYIR" {
		t.Errorf(`%s != "PAHNAPLSIIGYIR"`, ans)
	}

	if ans := convert("PAYPALISHIRING", 4); ans != "PINALSIGYAHRPI" {
		t.Errorf(`%s != "PINALSIGYAHRPI"`, ans)
	}

	if ans := convert("A", 1); ans != "A" {
		t.Errorf(`%s != "A"`, ans)
	}
}
