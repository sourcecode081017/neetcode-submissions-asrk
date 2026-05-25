/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }
    q := []*TreeNode{root}

    for len(q) > 0 {
        s := len(q)
        il := []int{}
        for i := 0; i < s; i++ {
            n := q[0]
            q = q[1:]
            il = append(il, n.Val)
            if n.Left != nil {
                q = append(q, n.Left)
            }
            if n.Right != nil {
                q = append(q, n.Right)
            }
            
        }
        res = append(res, il)
    }
    return res
}
