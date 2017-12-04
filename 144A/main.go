package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	min := 1000
	max := 0
	var minIdx, maxIdx int
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if a[i] <= min {
			min = a[i]
			minIdx = i
		}
		if a[i] > max {
			max = a[i]
			maxIdx = i
		}
	}
	if maxIdx > minIdx {
		fmt.Print(maxIdx + n - minIdx - 2)
		return
	}
	fmt.Print(maxIdx + n - minIdx - 1)
}
