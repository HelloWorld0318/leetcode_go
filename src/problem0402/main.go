package problem0402

func removeKdigits(num string, k int) string {
	stack := make([]int, 0)
	i, length := 0, len(num)
	for i < length {
		curNumber := int(num[i] - '0')
		for len(stack) != 0 && stack[len(stack)-1] > curNumber && k > 0 {
			stack = stack[0 : len(stack)-1]
			k--
		}
		//消除前导0
		if !(curNumber == 0 && len(stack) == 0) {
			stack = append(stack, curNumber)
		}
		i++
	}
	for k > 0 && len(stack) != 0 {
		k--
		stack = stack[0 : len(stack)-1]
	}
	var res string
	for _, temp := range stack {
		res += string('0' + temp)
	}
	if res == "" {
		res = "0"
	}
	return res
}
