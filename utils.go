package main

import (
	"strconv"
	"strings"
)

func compareSlices[T comparable](a, b []T) bool {
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

func compare2Dslices[T comparable](a, b [][]T) bool {
	n, m := len(a), len(a[0])

	for i := range n {
		for j := range m {
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
