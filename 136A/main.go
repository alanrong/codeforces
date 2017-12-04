package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+1)
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
		b[a[i]] = i
	}
	for i := 1; i <= n; i++ {
		fmt.Print(b[i])
		fmt.Print(" ")
	}
}
