package main

import "fmt"

/**
 * author: will fan
 * created: 2020/4/14 12:28
 * description:
 */
func main() {
	var x interface{} = "foo"

	var s string = x.(string)
	fmt.Println(s)

	s, ok := x.(string)
	fmt.Println(s, ok)

	n, ok := x.(int)
	fmt.Println(n, ok)

	//n = x.(int)

	switch v := x.(type) {
	case nil:
		fmt.Println("x is nil")
	case int:
		fmt.Println("x is ", v)
	case bool, string:
		fmt.Println("x is bool or string")
	default:
		fmt.Println("type unknown")
	}
}
