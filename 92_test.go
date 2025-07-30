package main

import "testing"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	sentinel := &ListNode{Next: head}
	sp := sentinel
	for range left - 1 {
		sp = sp.Next
	}

	var prev *ListNode
	curr := sp.Next
	for range right - left + 1 {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	sp.Next.Next, sp.Next = curr, prev

	return sentinel.Next
}

func TestReverseBetween(t *testing.T) {

	ans := reverseBetween(ToLinkedList(1, 2, 3, 4, 5), 2, 4)
	exp := ToLinkedList(1, 4, 3, 2, 5)
	if !ans.Compare(exp) {
		t.Errorf("TEST 1 FAILED: %s != %s", ans, exp)
	}

	ans = reverseBetween(ToLinkedList(5), 1, 1)
	exp = ToLinkedList(5)
	if !ans.Compare(exp) {
		t.Errorf("TEST 2 FAILED: %s != %s", ans, exp)
	}

	ans = reverseBetween(ToLinkedList(3, 5), 1, 1)
	exp = ToLinkedList(3, 5)
	if !ans.Compare(exp) {
		t.Errorf("TEST 3 FAILED: %s != %s", ans, exp)
	}
}
