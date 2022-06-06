package problem0224

func calculate(s string) int {
	var res, sign int
	sign = 1
	stack := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' {
			var num int
			for i < len(s) && s[i] > '0' {
				num = 10*num + int(s[i]-'0')
				i++
			}
			res += sign * num
			i--
		} else if s[i] == '+' {
			sign = 1
		} else if s[i] == '-' {
			sign = -1
		} else if s[i] == '(' {
			stack = append(stack, res)
			stack = append(stack, sign)
			res = 0
			sign = 1
		} else if s[i] == ')' {
			var signTemp, resTemp int
			signTemp, stack = stack[len(stack)-1], stack[:len(stack)-1]
			res *= signTemp
			resTemp, stack = stack[len(stack)-1], stack[:len(stack)-1]
			res += resTemp
		}
	}
	return res
}
