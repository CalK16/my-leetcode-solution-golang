package main

import (
	"slices"
	"testing"
)

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		for i := range len(a) {
			if a[i] > b[i] {
				return 1
			} else if a[i] < b[i] {
				return -1
			}
		}
		return 0
	})

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
	if !compare2Dslices(input, exp) {
		t.Errorf("TEST 1 FAILED")
	}
}
