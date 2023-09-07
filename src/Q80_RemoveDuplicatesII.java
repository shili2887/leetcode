package src;

public class Q80_RemoveDuplicatesII {

    /**
     * 使用双指针，一个指针指向当前不重复的元素的位置，另一个指针遍历整个数组，如果当前元素与前一个元素相同，则计数器加1，直到 2 停止。计数器给下一个
     * 不同的数字用。此时 i 指针不动，将当前元素复制到 i 指针指向的位置，并将 i 向后移动一位，否则不做任何操作。最后返回第一个指针指向的位置即可。
     * 时间复杂度：O(n)，其中n为数组的长度。
     *
     * @param nums 数组
     * @return int 数组最后长度
     */
    public static int removeDuplicates(int[] nums) {
        // [1,1,1,2,2,3]
        if (nums == null || nums.length == 0) {
            return 0;
        }
        int i = 1;
        int count = 1;
        for (int j = 1; j < nums.length; j++) {
            if (nums[j] == nums[j - 1]) {
                count++;
            } else {
                count = 1;
            }
            if (count <= 2) {
                nums[i] = nums[j];
                i++;
            }
        }
        return i;
    }

}
