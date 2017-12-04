package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	var a []string
	l := s
	var d int
	for {
		d = strings.Index(l, "WUB")
		if d == -1 {
			if len(l) > 0 {
				a = append(a, l)
			}
			break
		}
		c := l[:d]
		if len(c) > 0 {
			a = append(a, c)
		}
		if d+3 >= len(l) {
			l = ""
			break
		}
		l = l[d+3:]
	}
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i])
		fmt.Print(" ")
	}

}
