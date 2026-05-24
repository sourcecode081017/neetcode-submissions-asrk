/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    var dfs func(*TreeNode) int
	balanced := true
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lh := dfs(root.Left)
		rh := dfs(root.Right)
		if math.Abs(float64(lh) - float64(rh)) > 1.0 {
			balanced = false
			return -1
		}
		return 1 + max(lh, rh)
	}
	dfs(root)
	return balanced
}
