package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func main() {
	x, y := 10, 30

	result := add(x, y)
	fmt.Println("add:", result)

	result = subtract(x, y)
	fmt.Println("subtract:", result)
}
