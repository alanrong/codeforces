package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	var sum int
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		sum += a[i]
	}
	sort.Ints(a)
	var total int
	for i := n - 1; i >= 0; i-- {
		total += a[i]
		if total > sum/2 {
			fmt.Print(n - i)
			return
		}
	}
}
