package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+1)
	var p, q int
	fmt.Scan(&p)
	for i := 0; i < p; i++ {
		var j int
		fmt.Scan(&j)
		a[j]++
	}
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var j int
		fmt.Scan(&j)
		a[j]++
	}
	for i := 1; i <= n; i++ {
		if a[i] == 0 {
			fmt.Print("Oh, my keyboard!")
			return
		}
	}
	fmt.Print("I become the guy.")
}
