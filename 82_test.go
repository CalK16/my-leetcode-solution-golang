package main

import "testing"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	curr := dummy
	for curr.Next != nil && curr.Next.Next != nil {
		if curr.Next.Val == curr.Next.Next.Val {
			x := curr.Next.Val
			for curr.Next != nil && curr.Next.Val == x {
				curr.Next = curr.Next.Next
			}
		} else {
			curr = curr.Next
		}
	}

	return dummy.Next
}

func TestDeleteDuplicates(t *testing.T) {
	i := ToLinkedList(1, 2, 3, 3, 4, 4, 5)
	ans := deleteDuplicates(i)
	if !ans.Compare(ToLinkedList(1, 2, 5)) {
		t.Errorf("TEST 1 FAILED: %s", ans)
	}

	i = ToLinkedList(0, 0, 1, 1, 1, 2)
	ans = deleteDuplicates(i)
	if !ans.Compare(ToLinkedList(2)) {
		t.Errorf("TEST 1 FAILED: %s", ans)
	}
}
