package main

import "slices"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree106(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	n := len(postorder)
	splitAt := slices.Index(inorder, postorder[n-1])

	return &TreeNode{
		Val:   postorder[n-1],
		Left:  buildTree(inorder[:splitAt], postorder[:splitAt]),
		Right: buildTree(inorder[splitAt+1:], postorder[splitAt:n-1]),
	}
}
