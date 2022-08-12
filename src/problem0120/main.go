package problem0120

func minimumTotal(triangle [][]int) int {
	row, i := len(triangle), 0
	if row == 1 {
		res := triangle[0][0]
		for _, item := range triangle[0] {
			if item < res {
				res = item
			}
		}
		return res
	}
	dp, i := make([][]int, row), 0
	for i < len(triangle[row-1]) {
		dp[row-1] = append(dp[row-1], triangle[row-1][i])
		i++
	}
	i = row - 2
	for i >= 0 {
		j := 0
		dp[i] = make([]int, len(triangle[i]))
		for j < len(triangle[i]) {
			dp[i][j] = triangle[i][j] + min(dp[i+1][j], dp[i+1][j+1])
			j++
		}
		i--
	}
	return dp[0][0]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
