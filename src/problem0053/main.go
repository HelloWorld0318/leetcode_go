package problem0053

func maxSubArray(nums []int) int {
	//将求n个数的数组的最大子段和，转换成分别求以第1个，第2个，、、、、，第i个，、、、
	//第n个数字结尾的最大子段和，再找出这n个结果中最大的，即为结果
	//若dp[i-1]>0，dp[i] = dp[i-1] + nums[i]；否则：dp[i] = nums[i]
	//边界值：以第1个数字结尾的最大子段和dp[0]=nums[0]
	dp, i := make([]int, len(nums)), 0
	for i < len(nums) {
		if i == 0 {
			dp[0] = nums[0]
		} else {
			dp[i] = nums[i]
			if dp[i-1] > 0 {
				dp[i] = nums[i] + dp[i-1]
			}
		}
		i++
	}
	res, i := dp[0], 0
	for i < len(nums) {
		if dp[i] > res {
			res = dp[i]
		}
		i++
	}
	return res
}
