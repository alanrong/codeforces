package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	fmt.Print(strings.Compare(a, b))
}
