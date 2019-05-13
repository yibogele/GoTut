package main

import (
	"fmt"
)
/**
 * created: 2019/5/13 15:33
 * By Will Fan
 */
func main() {
	a := []int{1, 2}
	b := []int{2,3,4}
	a = append(a, b...)

	fmt.Println(a)

	fmt.Println(append(a, a...))
}
