package main

import (
	"fmt"
	"testing"
)

func squareSum(n int) int {
	ans := 0
	for n != 0 {
		ans += (n % 10) * (n % 10)
		n /= 10
	}
	return ans
}

func isHappy(n int) bool {
	slow, fast := n, squareSum(n)
	for slow != fast {
		slow = squareSum(n)
		fast = squareSum(squareSum(fast))
		fmt.Println(slow, fast)
	}
	return slow == 1
}

func TestIsHappy(t *testing.T) {
	if !isHappy(19) {
		t.Errorf("TEST 1 FAILED")
	}

	if isHappy(2) {
		t.Errorf("TEST 2 FAILED")
	}

	// if isHappy(2147483647) {

	// }
}
