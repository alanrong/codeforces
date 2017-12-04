package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	total := n
	p := n / m
	mod := n % m
	for p > 0 {
		total += p
		prev := mod
		mod = (p + prev) % m
		p = (p + prev) / m
	}
	fmt.Print(total)
}
