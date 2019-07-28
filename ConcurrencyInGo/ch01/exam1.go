package main

import "fmt"

/**
 * author: will fan
 * created: 2019/7/21 8:33
 * description:
 */
func main() {
	var data int
	go func() {
		data++
	}()

	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}
