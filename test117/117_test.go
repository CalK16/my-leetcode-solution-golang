package test117

import "testing"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	q := make([]*Node, 0)
	q = append(q, root)

	for len(q) > 0 {
		l := len(q)

		var prev *Node
		for i := range l {
			if prev != nil {
				prev.Next = q[i]
			}

			prev = q[i]

			if prev.Left != nil {
				q = append(q, prev.Left)
			}

			if prev.Right != nil {
				q = append(q, prev.Right)
			}
		}
		q = q[l:]
	}

	return root
}

func TestConnect(t *testing.T) {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n7 := &Node{Val: 7}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Right = n7

	_ = connect(n1)
	if n2.Next != n3 {
		t.Errorf("ERROR")
	}

	if n4.Next != n5 {
		t.Errorf("ERROR")
	}

	if n5.Next != n7 {
		t.Errorf("ERROR")
	}
}
