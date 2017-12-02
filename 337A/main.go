package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	min := 1000
	for i := 0; i < m-n+1; i++ {
		d := a[i+n-1] - a[i]
		if d < min {
			min = d
		}
	}
	fmt.Print(min)
}
