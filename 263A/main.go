package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var m, n int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&a)
			if a == 1 {
				m = i
				n = j
			}
		}
	}
	k := math.Abs(float64(m-2)) + math.Abs(float64(n-2))
	fmt.Print(int(k))
	return
}
