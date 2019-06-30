package main

import "fmt"

/**
 * author: will fan
 * created: 2019/6/30 17:14
 * description:
 */
func main() {
	x := []int {1, 2, 3, 4}
	for k, v := range x{
		fmt.Println(k, v)
	}
}
