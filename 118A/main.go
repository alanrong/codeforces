package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	str = strings.ToLower(str)
	var res string
	for i := 0; i < len(str); i++ {
		if isVow(string(str[i])) {
			continue
		}
		res += string(".") + string(str[i])
	}
	fmt.Print(res)
}

func isVow(b string) bool {
	if b == "a" ||
		b == "o" ||
		b == "y" ||
		b == "e" ||
		b == "u" ||
		b == "i" {
		return true
	}
	return false
}
