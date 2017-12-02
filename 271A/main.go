package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := n + 1; i < 10000; i++ {
		if isBeautiful(i) {
			fmt.Print(i)
			return
		}
	}
}

func isBeautiful(n int) bool {
	var a [10]int
	for n > 0 {
		i := n % 10
		a[i]++
		n = n / 10
	}
	for i := 0; i < 10; i++ {
		if a[i] >= 2 {
			return false
		}
	}
	return true
}
