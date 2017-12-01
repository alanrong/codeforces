package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	var arr []int
	for i := 0; i < len(s); i++ {
		if s[i] >= '1' && s[i] <= '3' {
			n, _ := strconv.Atoi(string(s[i]))
			arr = append(arr, n)
		}
	}
	sort.Ints(arr)
	for i := 0; i < len(arr)-1; i++ {
		fmt.Print(arr[i])
		fmt.Print("+")
	}
	fmt.Print(arr[len(arr)-1])
}
