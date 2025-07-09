package main

import (
	"strings"
	"testing"
)

var R2I = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

var Keys = []string{"CM", "CD", "XC", "XL", "IX", "IV", "M", "D", "C", "L", "X", "V", "I"}

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	for _, k := range Keys {
		if !strings.HasPrefix(s, k) {
			continue
		}

		_, tail, _ := strings.Cut(s, k)
		sumTail := romanToInt(tail)
		return R2I[k] + sumTail
	}
	return 0
}

func TestRomanToInt(t *testing.T) {
	if ans := romanToInt("III"); ans != 3 {
		t.Errorf("%d != 3", ans)
	}

	if ans := romanToInt("LVIII"); ans != 58 {
		t.Errorf("%d != 58", ans)
	}

	if ans := romanToInt("MCMXCIV"); ans != 1994 {
		t.Errorf("%d != 1994", ans)
	}
}
