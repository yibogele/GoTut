package main

import "fmt"

/**
 * author: will fan
 * created: 2019/6/30 17:12
 * description:
 */
func main() {
	var x []int
	fmt.Println(x, len(x), cap(x))

	x = append(x, 10, 20, 30)
	fmt.Println("Slice x:", x)
}
