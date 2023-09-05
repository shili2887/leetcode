package src;

public class Q45_JumpGameII {

    public int jump(int[] nums) {
        int n = nums.length;
        int maxPos = 0;
        int end = 0;
        int steps = 0;
        for (int i = 0; i < n - 1; i++) {
            maxPos = Math.max(maxPos, i + nums[i]); // 当前最远距离
            if (i == end) { // 如果遍历到当前需要跳跃的步数，需要更新步数并且更新当前需要跳跃的步数
                steps++;
                end = maxPos;
            }
        }
        return steps;
    }
}
