package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func (listNode *ListNode) Compare(other *ListNode) bool {
	if listNode == nil && other == nil {
		return true
	} else if listNode == nil || other == nil {
		return false
	}

	return listNode.Val == other.Val && listNode.Next.Compare(other.Next)
}

type LinkedListBuilder struct {
	dummy *ListNode
	tail  *ListNode
}

func NewLinkedListBuilder() *LinkedListBuilder {
	dummy := &ListNode{}
	return &LinkedListBuilder{
		dummy: dummy,
		tail:  dummy,
	}
}

func (l *LinkedListBuilder) With(ln *ListNode) *LinkedListBuilder {
	dummy := &ListNode{Next: ln}
	tail := dummy
	for tail.Next != nil {
		tail = tail.Next
	}
	return &LinkedListBuilder{
		dummy: dummy,
		tail:  tail,
	}
}

func (l *LinkedListBuilder) Append(num int) *LinkedListBuilder {
	node := &ListNode{Val: num}
	l.tail.Next = node
	l.tail = node
	return l
}

func (l *LinkedListBuilder) Extend(num ...int) *LinkedListBuilder {
	for _, n := range num {
		l.Append(n)
	}
	return l
}

func (l *LinkedListBuilder) Build() *ListNode {
	return l.dummy.Next
}
