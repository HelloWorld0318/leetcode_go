package main

import "fmt"

func main() {
	target := 9
	nums := []int{2, 7, 11, 15}
	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for index, v := range nums {
		value, ok := hash[target-v]
		if ok {
			return []int{value, index}
		}
		hash[v] = index
	}
	return []int{}
}
