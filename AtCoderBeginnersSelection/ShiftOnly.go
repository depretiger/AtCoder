package main

import (
	"fmt"
)

func main() {
	//Nを読み込む
	n := 0
	count := 0
	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	for {
		flag := 0
		for i := 0; i < n; i++ {
			if arr[i]%2 == 1 {
				flag = 1
			}
			arr[i] /= 2
		}
		if flag == 1 {
			break
		}
		count++
	}
	fmt.Println(count)
}
