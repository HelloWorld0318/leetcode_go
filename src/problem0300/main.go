package problem0300

//O(n^2)
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//dp[i]究竟应该怎么定义?以nums[i]结尾的最长递增子序列，和最大子序列和相似之处
	//第i个状态dp[i]代表以第i个元素结尾的最长上升子序列的长度
	dp := make([]int, len(nums))
	i := 0
	dp[0] = 1
	var res = dp[0]
	i, j := 1, 0
	for i < len(nums) {
		j = 0
		maxDpBeforeI := 0
		for j < i {
			if nums[i] > nums[j] && dp[j] > maxDpBeforeI {
				maxDpBeforeI = dp[j]
			}
			j++
		}
		dp[i] = maxDpBeforeI + 1
		if dp[i] > res {
			res = dp[i]
		}
		i++
	}
	return res
}

//O(n*log(n))
func lengthOfLISII(nums []int) int {
	return 0
}
