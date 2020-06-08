package main

import "fmt"

func main() {
	target := 9
	slice := []int{2, 7, 11, 15}
	result := twoSum(slice, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	result := make([]int, 0)
	hash := make(map[int]int)
	for k, v := range nums {
		hash[v] = k
	}
	for k, v := range nums {
		temp := target - v
		val, tag := hash[temp]
		if tag && val != k {
			//tag判断哈希表中是否存在temp键，val！=k避免重复取数
			result = append(result, k)
			result = append(result, val)
			return result
		}
	}
	return []int{}
}
