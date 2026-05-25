/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */

class Solution {
    public List<List<Integer>> levelOrder(TreeNode root) {
        if(root == null) {
            return new ArrayList<>();
        }
        Queue<TreeNode> q = new LinkedList<>();
        List<List<Integer>> res = new ArrayList<>();
        q.offer(root);
        while(!q.isEmpty()) {
            int l = q.size();
            List<Integer> il = new ArrayList<>();
            for(int i = 0; i < l; i++) {
                TreeNode n = q.poll();
                il.add(n.val);
                if(n.left != null) {
                    q.offer(n.left);
                }
                if(n.right != null) {
                    q.offer(n.right);
                }
            }
            res.add(il);
        }
        return res;
    }
}
