package main

import (
	"strconv"
	"strings"
)

func isSameSlices[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range len(a) {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func isSameShape[T comparable](a, b [][]T) bool {
	ra, rb := len(a), len(b)
	if ra != rb {
		return false
	}

	for i := range ra {
		if len(a[i]) != len(b[i]) {
			return false
		}
	}
	return true
}

// isSame2DSlice returns true if two 2D slices' cell
// are identical
func isSame2DSlice[T comparable](a, b [][]T) bool {
	if !isSameShape(a, b) {
		return false
	}
	n := len(a)

	for i := range n {
		for j := range len(a[i]) {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

func ToLinkedList(num ...int) *ListNode {
	builder := NewLinkedListBuilder()
	return builder.Extend(num...).Build()
}

func (l *ListNode) String() string {
	var b strings.Builder
	b.WriteByte('[')
	for l != nil {
		b.WriteString(strconv.Itoa(l.Val))
		if l.Next != nil {
			b.WriteByte(',')
		}
		l = l.Next
	}
	b.WriteByte(']')
	return b.String()
}
