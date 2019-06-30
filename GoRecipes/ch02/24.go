package main

import (
	"fmt"
)

/**
 * author: will fan
 * created: 2019/6/30 13:47
 * description:
 */

func Add(x, y int) (result int) {
	result = x + y
	return
}
func main() {
	fmt.Println("Add: ", Add(2,3))
}
