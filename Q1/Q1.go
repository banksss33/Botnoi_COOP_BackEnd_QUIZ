package main

import (
	"fmt"
	"strings"
)

func main() {
	const STAR string = "*"
	var count int

	fmt.Scan(&count)

	for i := 1; i <= count; i++ {
		fmt.Println(strings.Repeat(STAR, i))
	}

	for i := count - 1; i > 0; i-- {
		fmt.Println(strings.Repeat(STAR, i))
	}
}
