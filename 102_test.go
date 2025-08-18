package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Pair struct {
	Node  *TreeNode
	Level int
}

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	q := make([]*Pair, 0)
	q = append(q, &Pair{Node: root, Level: 0})

	for len(q) != 0 {
		curr := q[0]
		q = q[1:]

		if len(ans) < curr.Level+1 {
			ans = append(ans, make([]int, 0))
		}

		ans[curr.Level] = append(ans[curr.Level], curr.Node.Val)

		if curr.Node.Left != nil {
			q = append(q, &Pair{Node: curr.Node.Left, Level: curr.Level + 1})
		}

		if curr.Node.Right != nil {
			q = append(q, &Pair{Node: curr.Node.Right, Level: curr.Level + 1})
		}
	}

	return ans
}
