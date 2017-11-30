package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	if arr[k-1] > 0 {
		if k == n {
			fmt.Print(k)
			return
		}
		for i := k; i < n; i++ {
			if arr[i] != arr[i-1] {
				fmt.Print(i)
				return
			}
		}
		fmt.Print(n)
		return
	}
	for i := k - 1; i >= 0; i-- {
		if arr[i] > 0 {
			fmt.Print(i + 1)
			return
		}
	}
	fmt.Print(0)
}
