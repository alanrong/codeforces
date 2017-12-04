package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Print(4, " ", n-4)
		return
	}
	if n-9 < 9 {
		fmt.Print(n-9, " ", 9)
	} else {
		fmt.Print(9, " ", n-9)
	}
}
