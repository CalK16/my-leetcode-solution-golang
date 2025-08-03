package main

import (
	"strconv"
)

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func evalRPN(tokens []string) int {
	operators := map[string]func(a, b int) int{
		"+": add,
		"-": subtract,
		"*": multiply,
		"/": divide,
	}

	stack := make([]int, 0)
	for _, token := range tokens {
		if f, ok := operators[token]; ok {
			n := len(stack)
			a, b := stack[n-2], stack[n-1]
			stack = stack[:n-2]
			stack = append(stack, f(a, b))
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[len(stack)-1]
}
