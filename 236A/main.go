package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var a [26]int
	for i := 0; i < len(s); i++ {
		v := int(s[i] - 'a')
		a[v]++
	}
	var m int
	for i := 0; i < 26; i++ {
		if a[i] > 0 {
			m++
		}
	}
	if m%2 == 0 {
		fmt.Print("CHAT WITH HER!")
		return
	}
	fmt.Print("IGNORE HIM!")
}
