/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    return dfs(root, math.MinInt, math.MaxInt)
}
func dfs(root *TreeNode, lb, ub int) bool {
    if root == nil {
        return true
    }
    if root.Val <= lb || root.Val >= ub {
        return false
    }
    return dfs(root.Left, lb, root.Val) && dfs(root.Right, root.Val, ub)
}
