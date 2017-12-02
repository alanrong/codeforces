package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n, &t)
	var s string
	fmt.Scan(&s)
	a := []byte(s)
	for i := 0; i < t; i++ {
		for j := 0; j < n; j++ {
			if a[j] == 'B' && j < n-1 {
				if a[j+1] == 'G' {
					a[j+1] = 'B'
					a[j] = 'G'
					j++
				}
			}
		}
	}
	fmt.Print(string(a))
}
