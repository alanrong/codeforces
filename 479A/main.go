package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a != 1 && b != 1 && c != 1 {
		fmt.Print(a * b * c)
		return
	}
	var r []int
	r = append(r, (a+b)*c)
	r = append(r, a*(b+c))
	r = append(r, a+b+c)
	sort.Ints(r)
	fmt.Print(r[2])
}
