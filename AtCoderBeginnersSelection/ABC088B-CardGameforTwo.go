package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	alice, bob := 0, 0
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])

	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for i, n := range a {
		if i%2 == 0 {
			alice += n
		} else {
			bob += n
		}

	}
	fmt.Println(alice - bob)
}
