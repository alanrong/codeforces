package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	if len(s) != len(t) {
		fmt.Print("NO")
		return
	}
	l := len(s)
	for i := 0; i < len(s); i++ {
		if s[i] != t[l-i-1] {
			fmt.Print("NO")
			return
		}
	}
	fmt.Print("YES")
}
