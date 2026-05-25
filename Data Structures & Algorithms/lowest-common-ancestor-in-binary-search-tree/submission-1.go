func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    var lca *TreeNode
    
    // REMOVED root, p, q from the parameters to avoid shadowing!
    var dfs func(*TreeNode) bool 
    dfs = func(node *TreeNode) bool {
        if node == nil {
            return false
        }
        
        var isNodePOrQ int
        if node.Val == p.Val || node.Val == q.Val {
            isNodePOrQ = 1
        }
        
        var isLeftSubTree, isRightSubTree int
        if dfs(node.Left) {
            isLeftSubTree = 1
        }
        if dfs(node.Right) {
            isRightSubTree = 1
        }

        if (isNodePOrQ + isLeftSubTree + isRightSubTree) >= 2 {
            lca = node
        }
        
        return (isNodePOrQ + isLeftSubTree + isRightSubTree) > 0
    }
    
    dfs(root)
    return lca
}