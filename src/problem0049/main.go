package problem0049

import (
	"strconv"
)

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	sortedStr2List := make(map[string][]string)
	for _, str := range strs {
		freqT := [26]int{}
		for _, l := range str {
			freqT[int(l)-97]++
		}
		hashS := ""
		for _, f := range freqT {
			hashS = hashS + strconv.Itoa(f) + ","
		}
		hashS = hashS[:len(hashS)-1]
		if list, ok := sortedStr2List[hashS]; ok {
			list = append(list, str)
			sortedStr2List[hashS] = list
		} else {
			sortedStr2List[hashS] = append(make([]string, 0), str)
		}
	}
	for _, list := range sortedStr2List {
		res = append(res, list)
	}
	return res
}
