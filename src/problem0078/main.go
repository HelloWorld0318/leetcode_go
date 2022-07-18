package problem0078

var res [][]int
var items []int

func subsets(nums []int) [][]int {
	res, items = make([][]int, 0), make([]int, 0)
	res = append(res, make([]int, 0))
	dfs(0, nums)
	return res
}

func dfs(i int, nums []int) {
	if i >= len(nums) {
		return
	}
	items = append(items, nums[i])
	dst := make([]int, len(items))
	copy(dst, items)
	res = append(res, dst)
	dfs(i+1, nums)
	items = items[:len(items)-1]
	dfs(i+1, nums)
}
