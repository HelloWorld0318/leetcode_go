package problem0001

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	item2index := make(map[int]int)
	for index, item := range nums {
		delta := target - item
		if indexTemp, ok := item2index[delta]; ok {
			res[0] = indexTemp
			res[1] = index
			return res
		}
		item2index[item] = index
	}
	return res
}
