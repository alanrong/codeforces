package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)
	var cnt int
	r := bufio.NewScanner(os.Stdin)
	var prev string
	for r.Scan() {
		t := r.Text()
		if t != prev {
			prev = t
			cnt++
		}
	}
	fmt.Print(cnt)
}
