package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var x, y, z int
	var count, money int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &money)
	for i := a; i >= 0; i-- {
		x = money
		x -= 500 * i
		if x > 0 {
			for j := b; j >= 0; j-- {
				y = x
				y -= 100 * j
				if y > 0 {
					for k := c; k >= 0; k-- {
						z = y
						z -= 50 * k
						if z == 0 {
							count++
						}
					}
				}
				if y == 0 {
					count++
				}
			}
		}
		if x == 0 {
			count++
		}
	}
	fmt.Printf("%d", count)
}
