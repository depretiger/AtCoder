package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&d[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(d)))
	count := 0
	a := 101
	for _, b := range d {
		if a != b {
			a = b
			count++
		}
	}
	fmt.Print(count)
}
