package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := getLuckyNums()
	if isLucky(n) || isAlmostLucky(n, a) {
		fmt.Print("YES")
		return
	}
	fmt.Print("NO")
}

func getLuckyNums() []int {
	var a []int
	for i := 1; i <= 1000; i++ {
		if isLucky(i) {
			a = append(a, i)
		}
	}
	return a
}

func isLucky(n int) bool {
	for n > 0 {
		m := n % 10
		if m != 4 && m != 7 {
			return false
		}
		n = n / 10
	}
	return true
}

func isAlmostLucky(n int, a []int) bool {
	for i := 0; i < len(a); i++ {
		if n%a[i] == 0 {
			return true
		}
	}
	return false
}
