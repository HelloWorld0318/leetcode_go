package problem0322

func coinChange(coins []int, amount int) int {
	//https://leetcode.com/problems/coin-change/
	//假如数组coins=[1,2,5,7,10],在计算dp[i]时，dp[0],dp[1],dp[2],...,dp[i-1]都是已知的：
	//而金额i可由:
	//金额i可由:金额i-1与coins[0](1)组合；金额i-2可由coins[1](2)组合，
	//金额i-5可由coins[2](5)组合；金额i-7可由coins[3](7)组合
	//金额i-10可由coins[4](10)组合

	//状态i可由状态i-1、i-2、i-5、i-7、i-10，5个状态所转移，故
	//dp[i]=min(dp[i-1],dp[i-2],dp[i-5],dp[i-7],dp[i-10])+1

	//dp[i]表示金额为i的时候最小值
	dp := make([]int, amount+1)
	i, j := 0, 0
	for i < len(dp) {
		//初始化数据，所有金额都无法可达
		dp[i] = -1
		i++
	}
	dp[0] = 0
	i = 1
	for i < amount+1 {
		j = 0
		for j < len(coins) {
			index := i - coins[j]
			if index >= 0 && dp[index] != -1 {
				if dp[index]+1 < dp[i] || dp[i] == -1 {
					dp[i] = dp[index] + 1
				}
			}
			j++
		}
		i++
	}
	return dp[amount]
}
