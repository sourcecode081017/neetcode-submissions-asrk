/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	maxDia := 0
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lh := dfs(root.Left)
		rh := dfs(root.Right)
		maxDia = max(maxDia, lh + rh)
		return 1 + max(lh, rh)
	}
	dfs(root)
	return maxDia
}
