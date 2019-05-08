package main

import "fmt"

/**
 * created: 2019/5/8 9:05
 * By Will Fan
 */
func main() {
	data := "A\xfe\x02\xff\x04"

	for _, v := range data {
		fmt.Printf("%#x ", v)
	}

	fmt.Println()

	for _, v := range []byte(data) {
		fmt.Printf("%#x ", v)
	}
}
