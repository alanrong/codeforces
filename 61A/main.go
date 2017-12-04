package main

import "fmt"

func main() {
	var m, n string
	fmt.Scan(&m, &n)
	for i := 0; i < len(m); i++ {
		if m[i] == n[i] {
			fmt.Print("0")
		} else {
			fmt.Print("1")
		}
	}
}
