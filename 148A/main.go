package main

import "fmt"

func main() {
	var k, l, m, n, d int
	fmt.Scan(&k, &l, &m, &n, &d)
	if k == 1 || l == 1 || m == 1 || m == 1 {
		fmt.Print(d)
		return
	}
	var cnt int
	for i := 1; i <= d; i++ {
		if i%k == 0 ||
			i%l == 0 ||
			i%m == 0 ||
			i%n == 0 {
			cnt++
		}
	}
	fmt.Print(cnt)
}
