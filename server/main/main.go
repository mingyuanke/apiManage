package main

import (
	"fmt"
)

func main() {
	fmt.Print(fibonacci(3000, 1, 0, 1))
}
func fibonacci(max int, s1 int, s2 int, index int) int {
	if (index < max) {
		return fibonacci(max, s1 + s2, s1, index + 1);
	} else if (max == index ) {
		return s1 + s2
	} else {
		return 1
	}
}
