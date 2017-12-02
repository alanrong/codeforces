package main

import "fmt"
import "strings"

func main() {
	var s string
	fmt.Scan(&s)
	pos := strings.Index(s, "h")
	if pos == -1 {
		fmt.Print("NO")
		return
	}
	s = s[pos+1:]
	pos = strings.Index(s, "e")
	if pos == -1 {
		fmt.Print("NO")
		return
	}
	s = s[pos+1:]
	pos = strings.Index(s, "l")
	if pos == -1 {
		fmt.Print("NO")
		return
	}
	s = s[pos+1:]
	pos = strings.Index(s, "l")
	if pos == -1 {
		fmt.Print("NO")
		return
	}
	s = s[pos+1:]
	pos = strings.Index(s, "o")
	if pos == -1 {
		fmt.Print("NO")
		return
	}
	fmt.Print("YES")
}
