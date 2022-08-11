package problem0198

func rob(nums []int) int {
	dp, i := make([]int, len(nums)), 0
	for i < len(nums) {
		if i == 0 {
			dp[0] = nums[0]
		} else if i == 1 {
			dp[1] = nums[1]
			if nums[1] < nums[0] {
				dp[1] = nums[0]
			}
		} else {
			dp[i] = dp[i-1]
			if dp[i-2]+nums[i] > dp[i] {
				dp[i] = dp[i-2] + nums[i]
			}
		}
		i++
	}
	return dp[len(dp)-1]
}
