package main

import (
	"slices"
	"testing"
)

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	slices.Sort(nums)
	n := len(nums)

	for first := range n {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		third := n - 1
		target := -nums[first]

		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			for second < third && nums[second]+nums[third] > target {
				third--
			}

			if second == third {
				break
			}

			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return ans
}

func TestThreeSum(t *testing.T) {
	if ans := threeSum([]int{-1, 0, 1, 2, -1, -4}); !compare2Dslices(ans, [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}) {
		t.Errorf("test1 failed")
	}
}
