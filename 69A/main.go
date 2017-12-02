package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var x, y, z int
	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		x += a
		y += b
		z += c
	}
	if x == 0 && y == 0 && z == 0 {
		fmt.Print("YES")
		return
	}
	fmt.Print("NO")
}
