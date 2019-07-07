package main

import (
	"fmt"
)

func main() {
	var n, a, b int
	var sum int
	sumsum := 0
	fmt.Scanf("%d %d %d", &n, &a, &b)
	for i := 1; i <= n; i++ {
		sum = sumdig(i)
		if sum >= a && sum <= b {
			sumsum += i
		}
	}
	fmt.Printf("%d", sumsum)
}
func sumdig(num int) (sum int) {
	sum = 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return
}
