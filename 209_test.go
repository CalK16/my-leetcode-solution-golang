package main

import (
	"math"
	"testing"
)

func minSubArrayLen(target int, nums []int) int {
	l, r, sum, ans := 0, 0, 0, math.MaxInt
	for l < len(nums) {
		for sum <= target && r < len(nums) {
			sum += nums[r]
			r++
		}

		if sum >= target {
			ans = min(ans, r-l)
			sum -= nums[l]
			l++
			continue
		}

		break
	}

	if ans == math.MaxInt {
		return 0
	}

	return ans
}

func TestMinSubArrayLen(t *testing.T) {
	if ans := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}); ans != 2 {
		t.Errorf("TEST 1 FAILED")
	}

	if ans := minSubArrayLen(4, []int{1, 4, 4}); ans != 1 {
		t.Errorf("TEST 2 FAILED")
	}

	if ans := minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}); ans != 0 {
		t.Errorf("TEST 3 FAILED: %d != 0\n", ans)
	}
}
