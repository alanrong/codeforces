package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	var res int
	for i := 0; i < n; i++ {
		var line string
		fmt.Scan(&line)
		if strings.Contains(line, "++") {
			res++
		} else if strings.Contains(line, "--") {
			res--
		}
	}
	fmt.Print(res)
}
