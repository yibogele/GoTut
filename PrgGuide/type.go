package main

import "fmt"


/**
 * created: 2019/7/15 13:34
 * By Will Fan
 */
func main() {
	var x interface{} = []int{1, 2, 3}
	xType := fmt.Sprintf("%T", x)
	fmt.Println(xType)

	var y interface{} = 2.3
	switch t := y.(type) {
	case int:
		fmt.Println("int: ", t)
	case float64:
		fmt.Println("float64", t)
	default:
		fmt.Println("unknown")
	}
}
