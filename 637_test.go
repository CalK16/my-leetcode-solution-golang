package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type data struct{ sum, count int }

func averageOfLevels(root *TreeNode) []float64 {
	l := make([]data, 0)
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if level < len(l) {
			l[level].sum += node.Val
			l[level].count++
		} else {
			l = append(l, data{node.Val, 1})
		}

		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)

	average := make([]float64, len(l))
	for i, d := range l {
		average[i] = float64(d.sum) / float64(d.count)
	}
	return average
}
