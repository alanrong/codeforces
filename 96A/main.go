package main

import (
	"fmt"
	"strings"
)

func main() {
	var line string
	fmt.Scan(&line)
	if strings.Contains(line, "0000000") ||
		strings.Contains(line, "1111111") {
		fmt.Print("YES")
		return
	}
	fmt.Print("NO")
	return
}
