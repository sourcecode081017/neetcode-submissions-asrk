/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    gn := 0
    var dfs func(*TreeNode, int)
    running := math.MinInt
    dfs = func(root *TreeNode, running int) {
        if root == nil {
            return
        }
        if root.Val >= running {
            gn += 1
        }
        running = max(running, root.Val)
        dfs(root.Left, running)
        dfs(root.Right, running)
    }
    dfs(root, running)
    return gn
    
}
