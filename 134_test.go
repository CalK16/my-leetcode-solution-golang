package main

import "testing"

func canCompleteCircuit(gas, cost []int) int {
	total, start, curr := 0, 0, 0
	for i := range len(gas) {
		total += gas[i] - cost[i]
		curr += gas[i] - cost[i]
		if curr < 0 {
			start = i + 1
			curr = 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}

func TestCanCompleteCircuit(t *testing.T) {
	if val := canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}); val != 3 {
		t.Errorf("val != 3")
	}

	if val := canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}); val != -1 {
		t.Errorf("%d != -1", val)
	}

	if val := canCompleteCircuit([]int{5, 8, 2, 8}, []int{6, 5, 6, 6}); val != 3 {
		t.Errorf("%d != 3", val)
	}
}
