package problem0064

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	row, column := len(grid), len(grid[0])
	i, j := 0, 0
	for i < row {
		dp[i] = make([]int, column)
		i++
	}
	dp[0][0] = grid[0][0]
	i = 1
	for i < row {
		dp[i][0] = grid[i][0] + dp[i-1][0]
		i++
	}
	j = 1
	for j < column {
		dp[0][j] = grid[0][j] + dp[0][j-1]
		j++
	}
	i = 1
	for i < row {
		j = 1
		for j < column {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			j++
		}
		i++
	}
	return dp[row-1][column-1]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
