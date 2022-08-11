package problem0070

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	//第i阶的爬法数量 = 第i-1阶的爬法数量 + 第i-2阶的爬法数量
	f1, f2 := 1, 2
	for n-2 > 0 {
		f1 = f1 + f2
		f2, f1 = f1, f2
		n--
	}
	return f2
}
