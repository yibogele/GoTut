package main

import "fmt"

/**
 * created: 2019/5/8 9:13
 * By Will Fan
 */
func main() {
	m := map[string]int {"one":1, "two":2, "three":3, "four":4}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
