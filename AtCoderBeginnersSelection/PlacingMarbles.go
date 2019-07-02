package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// テキストを１行読み込む
func StrStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput = scanner.Text()

	stringInput = strings.TrimSpace(stringInput)
	return
}

func main() {
	count := 0
	s := StrStdin()
	for i := 0; i < len(s); i++ {
		if s[i] == 49 {
			count++
		}
	}
	fmt.Println(count)
}
