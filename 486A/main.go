package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)
	fmt.Print(fn(n))
}

func fn(n int64) int64 {
	if n == 1 {
		return -1
	}
	if n%2 == 0 {
		return n / 2
	}
	return -1 - n/2
}
