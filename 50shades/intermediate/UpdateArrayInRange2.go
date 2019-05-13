package main

import (
	"fmt"
)

/**
 * created: 2019/5/13 8:50
 * By Will Fan
 */
func main() {
	data := []*struct{num int} {{1}, {2}, {3}}

	for _, v := range data {
		v.num *= 10
	}

	fmt.Println("data", data[0],data[1], data[2])
}
