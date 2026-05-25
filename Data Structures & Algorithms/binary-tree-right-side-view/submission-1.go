/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }
    q := []*TreeNode{root}
    for len(q) > 0 {
        level := len(q)
        for i := 0; i < level; i++ {
            n := q[0]
            q = q[1:]
            if i == level - 1 {
                res = append(res, n.Val)
            }
            if n.Left != nil {
                q = append(q, n.Left)
            }
            if n.Right != nil {
                q = append(q, n.Right)
            }
        }
    }
    return res
}
