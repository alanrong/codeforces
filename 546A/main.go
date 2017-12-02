package main

import "fmt"

func main() {
	var k, n, w int
	fmt.Scan(&k, &n, &w)
	total := k * w * (w + 1) / 2
	if total <= n {
		fmt.Print(0)
		return
	}
	fmt.Print(total - n)
}
