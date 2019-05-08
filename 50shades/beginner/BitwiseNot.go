package main

import "fmt"

/**
 * created: 2019/5/8 9:28
 * By Will Fan
 */
func main() {
	// wrong
	//fmt.Println(~2)

	// right
	//fmt.Println(^2)
	var d uint8 = 2
	fmt.Printf("%08b", ^d)
}
