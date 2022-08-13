package problem0300

//O(n^2)
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//dp[i]究竟应该怎么定义?以nums[i]结尾的最长递增子序列，和最大子序列和相似之处
	//第i个状态dp[i]代表以第i个元素结尾的最长上升子序列的长度
	dp := make([]int, len(nums))
	i := 0
	dp[0] = 1
	var res = dp[0]
	i, j := 1, 0
	for i < len(nums) {
		j = 0
		maxDpBeforeI := 0
		for j < i {
			if nums[i] > nums[j] && dp[j] > maxDpBeforeI {
				maxDpBeforeI = dp[j]
			}
			j++
		}
		dp[i] = maxDpBeforeI + 1
		if dp[i] > res {
			res = dp[i]
		}
		i++
	}
	return res
}

//O(n*log(n))
func lengthOfLISII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	stack := make([]int, 0)
	stack = append(stack, nums[0])
	i := 1
	for i < len(nums) {
		if stack[len(stack)-1] < nums[i] {
			stack = append(stack, nums[i])
		} else {
			j := 0
			//可以用二分查找优化一下
			for j < len(stack) {
				if stack[j] >= nums[i] {
					stack[j] = nums[i]
					break
				}
				j++
			}
		}
		i++
	}
	return len(stack)
}

func lengthOfLISIII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	stack := make([]int, 0)
	stack = append(stack, nums[0])
	i := 1
	for i < len(nums) {
		if nums[i] > stack[len(stack)-1] {
			stack = append(stack, nums[i])
		} else {
			pos := binarySearch(stack, nums[i])
			stack[pos] = nums[i]
		}
		i++
	}
	return len(stack)
}

func binarySearch(nums []int, target int) (index int) {
	index = -1
	begin, end := 0, len(nums)-1
	for index == -1 {
		mid := (begin + end) / 2
		if nums[mid] == target {
			index = mid
		} else if nums[mid] > target {
			if mid == 0 || nums[mid-1] < target {
				index = mid
			}
			end = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] > target {
				index = mid + 1
			}
			begin = mid + 1
		}
	}
	return index
}
