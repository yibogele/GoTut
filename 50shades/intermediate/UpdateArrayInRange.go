package main

import (
	"fmt"
)

/**
 * created: 2019/5/13 8:43
 * By Will Fan
 */
func main() {
	data := []int{1,2,3}

	//for _, v := range data {
	//	v *= 10
	//}

	for i, _ := range data {
		data[i] *= 10
	}

	fmt.Println("data", data)
}
