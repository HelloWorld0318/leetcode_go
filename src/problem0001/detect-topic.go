package main

import "fmt"

func main() {
	name := "1234#百万红包"
	words := ([]rune)(name)
	var inStr string
	for i := 0; i < len(words); i++ {
		fmt.Println(string(words[i : i+1]))
		inStr = inStr + string(words[i:i+1]) + ","
	}
	fmt.Println(inStr)

	fmt.Println(string(words[4]))
	fmt.Println(len(name))
}