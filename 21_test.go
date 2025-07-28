package main

import "testing"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	sentinel := &ListNode{}
	curr := sentinel
	for list1 != nil || list2 != nil {
		if list1 == nil {
			curr.Next = list2
			list2 = list2.Next
		} else if list2 == nil {
			curr.Next = list1
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			curr.Next = list2
			list2 = list2.Next
		} else {
			curr.Next = list1
			list1 = list1.Next
		}
		curr = curr.Next
		curr.Next = nil
	}
	return sentinel.Next
}

func TestMergeTwoLists(t *testing.T) {
	ans := mergeTwoLists(ToLinkedList(1, 2, 4), ToLinkedList(1, 3, 4))
	exp := ToLinkedList(1, 1, 2, 3, 4, 4)
	if !ans.Compare(exp) {
		t.Errorf("TEST 1 FAILED: %s != %s", ans, exp)
	}

	ans = mergeTwoLists(ToLinkedList(), ToLinkedList())
	exp = nil
	if !ans.Compare(exp) {
		t.Errorf("TEST 2 FAILED: %s != %s", ans, exp)
	}

	ans = mergeTwoLists(nil, ToLinkedList(0))
	exp = ToLinkedList(0)
	if !ans.Compare(exp) {
		t.Errorf("TEST 3 FAILED: %s != %s", ans, exp)
	}
}
