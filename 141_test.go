package main

import "testing"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	fast, slow := head, head.Next
	for fast != slow {
		fast = fast.Next
		if slow.Next == nil || slow.Next.Next == nil {
			return false
		}
		slow = slow.Next.Next
	}

	return true
}

func TestHasCycle(t *testing.T) {
	builder := NewLinkedListBuilder()
	linkedList := builder.Extend(3, 2, 0, -1).Build()
	linkedList.Next.Next.Next.Next = linkedList.Next
	if !hasCycle(linkedList) {
		t.Errorf("TEST 1 FAILED")
	}

	builder = NewLinkedListBuilder()
	linkedList = builder.Extend(1, 2).Build()
	linkedList.Next.Next = linkedList.Next
	if !hasCycle(linkedList) {
		t.Errorf("TEST 2 FAILED")
	}

	builder = NewLinkedListBuilder()
	linkedList = builder.Extend(1).Build()
	linkedList.Next = linkedList
	if !hasCycle(linkedList) {
		t.Errorf("TEST 3 FAILED")
	}

	builder = NewLinkedListBuilder()
	linkedList = builder.Build()
	if hasCycle(linkedList) {
		t.Errorf("TEST 4 FAILED")
	}
}
