package main

import (
	"fmt"
)
/**
 * created: 2019/5/13 15:36
 * By Will Fan
 */
func main() {
	var s = make([]int, 3)
	copy(s, []int{1,2,3,4})
	fmt.Println(s)

	s = []int{0,1,2}
	copy(s, s[1:])
	fmt.Println(s)

	var b = make([]byte, 5)
	copy(b, "Hello, World!")
	fmt.Println(b)

}
