package problem0376

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	maxLength := 1
	state, begin, up, down := 0, 0, 1, 2
	i := 1
	for i < len(nums) {
		switch state {
		case begin:
			if nums[i] > nums[i-1] {
				state = up
				maxLength++
			} else if nums[i] < nums[i-1] {
				state = down
				maxLength++
			}
			break
		case up:
			if nums[i] < nums[i-1] {
				state = down
				maxLength++
			}
			break
		case down:
			if nums[i] > nums[i-1] {
				state = up
				maxLength++
			}
			break
		}
		i++
	}
	return maxLength
}
