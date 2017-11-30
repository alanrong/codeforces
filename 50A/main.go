package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	if m%2 == 0 || n%2 == 0 {
		fmt.Print(m * n / 2)
		return
	}
	fmt.Print((m*n - 1) / 2)
	return
}
