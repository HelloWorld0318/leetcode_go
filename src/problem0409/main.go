package problem0409

func longestPalindrome(s string) int {
	charCounter := make(map[int]int)
	for _, char := range s {
		if count, ok := charCounter[int(char-'0')]; ok {
			charCounter[int(char-'0')] = count + 1
		} else {
			charCounter[int(char-'0')] = 1
		}
	}
	maxLength, flag := 0, 0
	for _, count := range charCounter {
		if count%2 == 0 {
			maxLength += count
		} else {
			maxLength += count - 1
			flag = 1
		}
	}
	return maxLength + flag
}
