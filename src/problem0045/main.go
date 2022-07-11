package problem0045

func jump(nums []int) int {
	//贪心规律：在到达某点前若一直不跳跃，发现从该点不能跳到更远的地方了，在这之前肯定有次必要的跳跃
	//结论：在无法到达更远的地方时，在这之前应该跳到一个可以到达更远的位置!

	//如果数组长度小于2，则说明不用跳跃，返回0
	if len(nums) < 2 {
		return 0
	}
	//当前可达到的最远位置；遍历各个位置过程中，可达到的最远位置
	currentMaxIndex, preMaxIndex, jumpMin := nums[0], nums[0], 1
	i := 1
	for i < len(nums) {
		//当无法再向前移动了，才进行跳跃
		if i > currentMaxIndex {
			jumpMin++
			//更新到当前可达到的最远位置
			currentMaxIndex = preMaxIndex
		}
		if preMaxIndex < nums[i]+i {
			preMaxIndex = nums[i] + i
		}
		i++
	}
	return jumpMin
}
