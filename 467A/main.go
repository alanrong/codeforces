package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var cnt int
	for i := 0; i < n; i++ {
		var m, l int
		fmt.Scan(&m, &l)
		if l-m >= 2 {
			cnt++
		}
	}
	fmt.Print(cnt)
}
