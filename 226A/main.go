package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	var cnt int
	c := s[0]
	for i := 1; i < n; i++ {
		if c == s[i] {
			cnt++
		} else {
			c = s[i]
		}
	}
	fmt.Print(cnt)
}
