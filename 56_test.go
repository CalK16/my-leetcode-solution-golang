package main

import (
	"slices"
	"testing"
)

func merge(intervals [][]int) [][]int {
	ret := make([][]int, 0)
	slices.SortFunc(intervals, func(a, b []int) int { return a[0] - b[0] })

	for _, interval := range intervals {
		m := len(ret)
		if m > 0 && interval[0] <= ret[m-1][1] {
			ret[m-1][1] = max(ret[m-1][1], interval[1])
		} else {
			ret = append(ret, interval)
		}
	}

	return ret
}

func TestMerge(t *testing.T) {
	input := [][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}
	exp := [][]int{
		[]int{1, 6},
		[]int{8, 10},
		[]int{15, 18},
	}
	if !isSame2DSlice(merge(input), exp) {
		t.Errorf("TEST 1 FAILED: %v != %v", input, exp)
	}
}
