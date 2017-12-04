package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var max int
	cur := 1
	for i := 1; i < n; i++ {
		if a[i] >= a[i-1] {
			cur++
		} else {
			if cur > max {
				max = cur
			}
			cur = 1
		}
	}
	if cur > max {
		max = cur
	}
	fmt.Print(max)
}
