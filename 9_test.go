package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	d := make([]int, 0)
	for x != 0 {
		d = append(d, x%10)
		x /= 10
	}

	start, end := 0, len(d)-1
	for start <= end {
		if d[start] != d[end] {
			return false
		}

		start++
		end--
	}

	return true
}
