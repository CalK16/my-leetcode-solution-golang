package main

import "testing"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	carryFlag := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carryFlag
		curr.Next = &ListNode{
			Val: sum % 10,
		}
		curr = curr.Next
		carryFlag = sum / 10
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum := l1.Val + carryFlag
		curr.Next = &ListNode{
			Val: sum % 10,
		}
		curr = curr.Next
		carryFlag = sum / 10
		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val + carryFlag
		curr.Next = &ListNode{
			Val: sum % 10,
		}
		curr = curr.Next
		carryFlag = sum / 10
		l2 = l2.Next
	}

	if carryFlag != 0 {
		curr.Next = &ListNode{Val: 1}
	}

	return dummy.Next
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := ToLinkedList(2, 4, 3)
	l2 := ToLinkedList(5, 6, 4)

	if !addTwoNumbers(l1, l2).Compare(ToLinkedList(7, 0, 8)) {
		t.Errorf("TEST 1 FAILED")
	}

	l1 = ToLinkedList(9, 9, 9, 9, 9, 9, 9)
	l2 = ToLinkedList(9, 9, 9, 9)
	if !addTwoNumbers(l1, l2).Compare(ToLinkedList(8, 9, 9, 9, 0, 0, 0, 1)) {
		t.Errorf("TEST 2 FAILED: %s != %s", addTwoNumbers(l1, l2), ToLinkedList(8, 9, 9, 9, 0, 0, 0, 1))
	}
}
