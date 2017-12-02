package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var cur, max int
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		cur += b - a
		if cur > max {
			max = cur
		}
	}
	fmt.Print(max)
}
