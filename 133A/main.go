package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		if s[i] == 'H' || s[i] == '9' || s[i] == 'Q' {
			fmt.Print("YES")
			return
		}
	}
	fmt.Print("NO")
}
