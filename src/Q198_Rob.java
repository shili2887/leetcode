package src;

public class Q198_Rob {

    /**
     * 1. 偷：如果偷第 i 个房屋，则不能偷第 i-1 个房屋，因此能够获得的最大金额为 dp[i-2] + nums[i]。
     * 2. 不偷：如果不偷第 i 个房屋，则能够获得的最大金额为 dp[i-1]。
     * 状态转移方程为：dp[i] = max(dp[i-2]+nums[i], dp[i-1])。
     *
     * @param nums 数组
     * @return 最大金额
     */
    public int rob(int[] nums) {
        int n = nums.length;
        if (n == 0) {
            return 0;
        }
        if (n == 1) {
            return nums[0];
        }
        int[] dp = new int[n];
        dp[0] = nums[0];
        dp[1] = Math.max(nums[0], nums[1]);
        for (int i = 2; i < n; i++) {
            dp[i] = Math.max(dp[i - 2] + nums[i], dp[i - 1]);
        }
        return dp[n - 1];
    }

}
