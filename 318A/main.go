package main

import "fmt"

func main() {
	var n, k int64
	fmt.Scan(&n, &k)
	if n%2 == 0 {
		if k <= n/2 {
			fmt.Print(k*2 - 1)
		} else {
			fmt.Print((k - n/2) * 2)
		}
		return
	}
	if n%2 == 1 {
		if k <= n/2+1 {
			fmt.Print(k*2 - 1)
		} else {
			fmt.Print((k - n/2 - 1) * 2)
		}
		return
	}
}
