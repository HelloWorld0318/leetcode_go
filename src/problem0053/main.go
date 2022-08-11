package problem0053

func maxSubArray(nums []int) int {
	dp, i := make([]int, len(nums)), 0
	for i < len(nums) {
		if i == 0 {
			dp[0] = nums[0]
		} else {
			dp[i] = nums[i]
			if nums[i]+dp[i-1] > nums[i] {
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
