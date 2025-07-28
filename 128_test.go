package main

import "testing"

func longestConsecutive(nums []int) int {
	ans := 0
	m := make(map[int]bool)
	for _, x := range nums {
		m[x] = true
	}

	for x, _ := range m {
		if _, found := m[x-1]; found {
			continue
		}

		y := x + 1
		for m[y] {
			y++
		}

		ans = max(ans, y-x)
	}

	return ans
}

func TestLongestConsecutive(t *testing.T) {
	if longestConsecutive([]int{100, 4, 200, 1, 3, 2}) != 4 {
		t.Errorf("TEST 1 FAILED")
	}

	if longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}) != 9 {
		t.Errorf("TEST 2 FAILED")
	}

	if longestConsecutive([]int{1, 0, 1, 2}) != 3 {
		t.Errorf("TEST 3 FAILED")
	}
}
