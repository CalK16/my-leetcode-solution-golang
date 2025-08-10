package main

import (
	"slices"
	"testing"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	if len(intervals) == 1 {
		return intervals
	}
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	ans := intervals[:1]
	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		if curr[0] > ans[len(ans)-1][1] {
			ans = append(ans, curr)
		} else {
			ans[len(ans)-1][1] = max(curr[1], ans[len(ans)-1][1])
		}
	}
	return ans
}

func TestInsert(t *testing.T) {
	interval := [][]int{
		[]int{1, 3},
		[]int{6, 9},
	}
	newInterval := []int{2, 5}
	exp := [][]int{
		[]int{1, 5},
		[]int{6, 9},
	}

	ans := insert(interval, newInterval)
	if !isSame2DSlice(ans, exp) {
		t.Errorf("TEST 1 FAILED")
	}

}
