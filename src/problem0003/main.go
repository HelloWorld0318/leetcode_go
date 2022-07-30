package problem0003

func lengthOfLongestSubstring(s string) int {
	start, res, char2Index := -1, 0, make(map[string]int)
	for index, char := range s {
		if charIndex, ok := char2Index[string(char)]; ok && charIndex > start {
			start = charIndex
		}
		char2Index[string(char)] = index
		if index-start > res {
			res = index - start
		}
	}
	return res
}
