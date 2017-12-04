package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%5 == 0 {
		fmt.Print(n / 5)
		return
	}
	fmt.Print(n/5 + 1)
}
