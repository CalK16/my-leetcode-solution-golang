package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	tail := head
	for range n {
		tail = tail.Next
	}
	prev := dummy
	for tail != nil {
		prev = prev.Next
		tail = tail.Next
	}
	prev.Next = prev.Next.Next
	return dummy.Next
}
