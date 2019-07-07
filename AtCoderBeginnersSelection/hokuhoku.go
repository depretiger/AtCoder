package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	s := ""
	rs := []rune("ホイクモ")
	for i, j := range rand.Perm(4)[:2] {
		s += string(rs[j])
		fmt.Println(i, j)
	}
	s += s + "の"
	for _, j := range rand.Perm(4)[:2] {
		s += string(rs[j])
	}
	fmt.Println(s)
}
