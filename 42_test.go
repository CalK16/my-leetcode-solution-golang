package main

import (
	"testing"
)

// solution 1
// Time: O(n)
// Space: O(n)
func trap(height []int) int {
	n := len(height)
	l2r := make([]int, n)

	for i := range n {
		l2r[i] = 1
	}

	localMax := 0
	for i := range n {
		localMax = max(localMax, height[i])
		l2r[i] = localMax
	}

	sum := 0
	localMax = 0
	for i := n - 1; i >= 0; i-- {
		localMax = max(localMax, height[i])
		sum += min(l2r[i], localMax) - height[i]
	}

	return sum
}

func TestTrap(t *testing.T) {
	if ans := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}); ans != 6 {
		t.Errorf("%d != 6", ans)
	}

	if ans := trap([]int{4, 2, 0, 3, 2, 5}); ans != 9 {
		t.Errorf("%d != 9", ans)
	}
}
