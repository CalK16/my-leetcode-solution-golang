package main

import "testing"

func candy(ratings []int) int {
	n := len(ratings)
	res := make([]int, n)

	for i := range n {
		res[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			res[i] = res[i-1] + 1
		}
	}

	sum := res[n-1]
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			res[i] = max(res[i], res[i+1]+1)
		}
		sum += res[i]
	}

	return sum
}

func TestCandy(t *testing.T) {
	if ans := candy([]int{1, 0, 2}); ans != 5 {
		t.Errorf("%d != 5", ans)
	}

	if ans := candy([]int{1, 2, 2}); ans != 4 {
		t.Errorf("%d != 4", ans)
	}
}
