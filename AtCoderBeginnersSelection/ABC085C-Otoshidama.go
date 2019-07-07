package main

import (
	"fmt"
)

func main() {
	var n, y int
	fmt.Scan(&n)
	fmt.Scan(&y)
	for a := 0; a <= n; a++ {
		for b := 0; b <= n-a; b++ {
			if y == 10000*a+5000*b+1000*(n-a-b) {
				fmt.Printf("%d %d %d", a, b, n-a-b)
				return
			}
		}
	}
	fmt.Printf("-1 -1 -1")
}
