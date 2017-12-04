package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	min := n
	if n > m {
		min = m
	}
	if min%2 == 0 {
		fmt.Print("Malvika")
		return
	}

	fmt.Print("Akshat")
}
