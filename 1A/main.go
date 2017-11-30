package main

import "fmt"

func main() {
	var n, m, a int64
	fmt.Scan(&n, &m, &a)
	var i, j int64
	if n%a == 0 {
		i = n / a
	} else {
		i = n/a + 1
	}

	if m%a == 0 {
		j = m / a
	} else {
		j = m/a + 1
	}
	fmt.Print(i * j)
}
