package src;

import java.util.Stack;

public class Q230_KthSmallest {

    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
    }

    public int kthSmallest(TreeNode root, int k) {
        Stack<TreeNode> stack = new Stack<>();
        TreeNode curr = root;
        int count = 0;
        while (curr != null || !stack.isEmpty()) {
            while (curr != null) {
                stack.push(curr);
                curr = curr.left;
            }
            curr = stack.pop();
            count++;
            if (count == k) {
                return curr.val;
            }
            curr = curr.right;
        }
        return -1;
    }

}
