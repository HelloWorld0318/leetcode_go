package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const _timeStamp20170701 = 1498838400

func main() {
	target := 9
	nums := []int{2, 7, 11, 15}
	result := twoSum(nums, target)
	fmt.Println(result)

	fmt.Println("============")

	shardInfo, _ := shardDynId(433238483638707017)
	fmt.Println("year:",shardInfo.Year)
	fmt.Println("month:",shardInfo.Month)
	fmt.Println("rand:",shardInfo.Rand)

	fmt.Println("============")
	blockUsers := make(map[int64]ExpiredTs)
	blockUsers[0] = ExpiredTs{ExpiredTs: 0}
	blockUsersBys, _ := json.Marshal(blockUsers)
	fmt.Println(string(blockUsersBys))

	fmt.Println(len("bk:1:212295231"))
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

func shardDynId(dynamicID uint64) (shardInfo *ShardInfo, err error) {
	ts := (dynamicID >> 32) + _timeStamp20170701
	tm := time.Unix(int64(ts), 0)
	fmt.Println("timestamp:", tm)
	shardInfo = &ShardInfo{
		Year:  int32(tm.Year()),
		Month: int32(tm.Month()),
		Rand:  (dynamicID & 0xffffffff) >> 20,
	}
	return
}

type ShardInfo struct {
	Year  int32
	Month int32
	Rand  uint64
}

type ExpiredTs struct {
	ExpiredTs int64 `json:"ex_ts"`
}
