package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	copied := make(map[*Node]*Node)
	curr := head
	for curr != nil {
		if copied[curr] == nil {
			copied[curr] = &Node{Val: curr.Val}
		}
		node := copied[curr]

		if curr.Next != nil && copied[curr.Next] == nil {
			copied[curr.Next] = &Node{Val: curr.Next.Val}
		}
		next := copied[curr.Next]
		node.Next = next

		if curr.Random != nil && copied[curr.Random] == nil {
			copied[curr.Random] = &Node{Val: curr.Random.Val}
		}
		random := copied[curr.Random]
		node.Random = random

		curr = curr.Next
	}

	return copied[head]
}
