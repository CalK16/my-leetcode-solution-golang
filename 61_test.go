package main

import "testing"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	tail := head
	n := 1
	for tail.Next != nil {
		tail = tail.Next
		n++
	}

	bp := head
	bi := n - k%n
	if bi == n {
		return head
	}

	for range bi - 1 {
		bp = bp.Next
	}

	newHead := bp.Next
	bp.Next = nil
	tail.Next = head
	return newHead
}

func TestRotateRight(t *testing.T) {
	ans := rotateRight(ToLinkedList(1, 2, 3, 4, 5), 2)
	exp := ToLinkedList(4, 5, 1, 2, 3)
	if !ans.Compare(exp) {
		t.Errorf("TEST 1 FAILED %s != %s", ans, exp)
	}

	ans = rotateRight(ToLinkedList(0, 1, 2), 4)
	exp = ToLinkedList(2, 0, 1)
	if !ans.Compare(exp) {
		t.Errorf("TEST 2 FAILED %s != %s", ans, exp)
	}
}
