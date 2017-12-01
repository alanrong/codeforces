package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var arr [5]int
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		arr[a]++
	}
	total := arr[4]
	total += arr[3]
	arr[1] -= arr[3]
	if arr[2]%2 == 0 {
		total += arr[2] / 2
	} else {
		total += arr[2]/2 + 1
		arr[1] -= 2
	}

	if arr[1] > 0 {
		if arr[1]%4 == 0 {
			total += arr[1] / 4
		} else {
			total += arr[1]/4 + 1
		}
	}
	fmt.Print(total)
}
