package problem0174

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 {
		return 0
	}
	row, column := len(dungeon), len(dungeon[0])
	i, j := 0, 0
	dp := make([][]int, row)
	for i < row {
		dp[i] = make([]int, column)
		i++
	}
	dp[row-1][column-1] = max(1, 1-dungeon[row-1][column-1])
	i = row - 2
	for i >= 0 {
		dp[i][column-1] = max(1, dp[i+1][column-1]-dungeon[i][column-1])
		i--
	}
	j = column - 2
	for j >= 0 {
		dp[row-1][j] = max(1, dp[row-1][j+1]-dungeon[row-1][j])
		j--
	}
	i = row - 2
	for i >= 0 {
		j = column - 2
		for j >= 0 {
			dpMin := min(dp[i+1][j], dp[i][j+1])
			dp[i][j] = max(1, dpMin-dungeon[i][j])
			j--
		}
		i--
	}
	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
