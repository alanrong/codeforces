package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var res []string
	for i := 0; i < n; i++ {
		var w string
		fmt.Scan(&w)
		w = short(w)
		res = append(res, w)
	}
	for _, v := range res {
		fmt.Println(v)
	}
}

func short(w string) string {
	if len(w) <= 10 {
		return w
	}
	l := len(w)
	return fmt.Sprintf("%c%d%c", w[0], l-2, w[l-1])
}
