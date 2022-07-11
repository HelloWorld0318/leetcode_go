package problem0055

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	jumpMaxIndexArr := make([]int, 0)
	for i, maxJump := range nums {
		jumpMaxIndexArr = append(jumpMaxIndexArr, i+maxJump)
	}

	jump := 0
	curMaxJumpIndex := nums[0]

	for jump <= curMaxJumpIndex && jump < len(nums) {
		if curMaxJumpIndex < jumpMaxIndexArr[jump] {
			curMaxJumpIndex = jumpMaxIndexArr[jump]
		}
		jump++
	}
	return jump == len(nums)
}
