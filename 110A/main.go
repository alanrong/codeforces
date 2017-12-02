package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)
	if isNearlyLucky(n) {
		fmt.Print("YES")
		return
	}
	fmt.Print("NO")
}

func isLucky(n int64) bool {
	if n == 0 {
		return false
	}
	for n > 0 {
		a := n % 10
		if a != 4 && a != 7 {
			return false
		}
		n = n / 10
	}
	return true
}

func isNearlyLucky(n int64) bool {
	var cnt int64
	for n > 0 {
		a := n % 10
		if a == 4 || a == 7 {
			cnt++
		}
		n = n / 10
	}
	if isLucky(cnt) {
		return true
	}
	return false
}
