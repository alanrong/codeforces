package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)
	if (input%2) == 0 && input > 2 {
		fmt.Print("YES")
		return
	}
	fmt.Print("NO")
	return
}
