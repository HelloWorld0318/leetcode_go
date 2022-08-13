package problem0120

func minimumTotal(triangle [][]int) int {
	//1.从上到下或者从下到上的寻找路径的思考方式本质是一样的吗？
	//2.假设dp[i][j]代表了数组三角形第i行，第j列的最优解，从上到下与从下到上的哪种方式递推更容易？（更少的边界调节）

	//从下到上的思考方式
	//1.设置一个二维数组，最优值三角形dp[][]，并初始化数组元素为0。dp[i][j]代表了从底向上推时，走到三角形第i行第j列的最优解。
	//2.从三角形的底部向三角形上方进行动态规划
	// 2.1动态规划边界条件：底部上的最优值即为数字三角形的最后一层。
	// 2.2利用i循环，从倒数第二层递推至第一层，对于每层的各列，进行动态规划递推：
	// 第i行，第j列的最优解为dp[i][j]，可到达(i,j)的两个位置的最优解dp[i+1][j]、dp[i+1][j+1]:
	// dp[i][j] = min(dp[i+1][j],dp[i+1][j+1]) + triangle[i][j]
	//3.返回dp[0][0]
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
