package problem0290

import "strings"

func wordPattern1(pattern string, s string) bool {
	strArr := strings.Split(s, " ")
	if len(pattern) != len(strArr) {
		return false
	}
	i := 0
	pattern2Str, str2Pattern := make(map[string]string), make(map[string]string)
	for i < len(pattern) {
		char := string(pattern[i])
		temp1, ok1 := pattern2Str[char]

		str := strArr[i]
		temp2, ok2 := str2Pattern[str]
		if (ok1 && temp1 != str) || (ok2 && char != temp2) {
			return false
		} else {
			pattern2Str[char] = str
			str2Pattern[str] = char
		}
		i++
	}
	return true
}

func wordPattern(pattern string, s string) bool {
	//单词到pattern字符的映射
	wordMap := make(map[string]string)
	//保存已被映射的pattern字符
	used := make([]int, 128)
	//保存临时拆分出的单词
	var word string
	//str尾部push一个空格，使得遇到空格拆分单词
	s += " "
	//pos 当前指向的pattern字符
	var pos, i int
	for i < len(s) {
		//遇到空格，即拆分出一个单词
		if s[i] == ' ' {
			//若分割出一个单词，但已无pattern字符对应
			if pos == len(pattern) {
				return false
			}
			if _, ok := wordMap[word]; ok {
				//若当前word已建立映射无法与当前pattern对应
				if wordMap[word] != string(pattern[pos]) {
					return false
				}
			} else {
				//若单词未出现在哈希映射中

				//如果当前pattern字符已经使用了
				if used[int(pattern[pos])] > 0 {
					return false
				}
				wordMap[word] = string(pattern[pos])
				used[int(pattern[pos])] = 1
			}
			pos++
			word = ""
		} else {
			word += string(s[i])
		}
		i++
	}
	if pos != len(pattern) {
		return false
	}
	return true
}
