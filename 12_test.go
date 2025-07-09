package main

import (
	"strings"
	"testing"
)

var I2R = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

var keys = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func intToRoman(num int) string {
	var ans strings.Builder

	for _, k := range keys {
		if k > num {
			continue
		}

		for range num / k {
			ans.WriteString(I2R[k])
		}

		num = num % k
	}

	return ans.String()
}

func TestIntToRoman(t *testing.T) {
	if ans := intToRoman(3749); ans != "MMMDCCXLIX" {
		t.Errorf("%s != MMMDCCXLIX", ans)
	}

	if ans := intToRoman(58); ans != "LVIII" {
		t.Errorf("%s != LVIII", ans)
	}

	if ans := intToRoman(1994); ans != "MCMXCIV" {
		t.Errorf("%s != MXMXCIV", ans)
	}
}
