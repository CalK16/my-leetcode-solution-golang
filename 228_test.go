package main

import (
	"fmt"
	"testing"
)

func summaryRanges(nums []int) []string {
	ans := make([]string, 0)
	ptr := 0
	for ptr < len(nums) {
		curr := ptr
		for curr+1 < len(nums) && nums[curr+1] == nums[curr]+1 {
			curr++
		}
		var item string
		if ptr == curr {
			item = fmt.Sprintf("%d", nums[ptr])
		} else {
			item = fmt.Sprintf("%d->%d", nums[ptr], nums[curr])
		}
		ans = append(ans, item)
		ptr = curr + 1
	}
	return ans
}

func TestSummaryRanges(t *testing.T) {
	ans := summaryRanges([]int{0, 1, 2, 4, 5, 7})
	exp := []string{"0->2", "4->5", "7"}
	if !compareSlices(ans, exp) {
		t.Errorf("TEST 1 FAILED: %s != %s", ans, exp)
	}

	ans = summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})
	exp = []string{"0", "2->4", "6", "8->9"}
	if !compareSlices(ans, exp) {
		t.Errorf("TEST 1 FAILED: %s != %s", ans, exp)
	}
}
