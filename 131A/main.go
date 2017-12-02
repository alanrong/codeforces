package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	var flag bool
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			flag = true
			break
		}
	}
	if flag {
		s = strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
	}
	fmt.Print(s)
}
