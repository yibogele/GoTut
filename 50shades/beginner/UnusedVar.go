package main

import "fmt"

var gvar int // not an error

/**
 * author: will fan
 * created: 2019/5/7 21:22
 * description:
 */
func main() {
	var one int // error
	two := 2	// error
	var three int	// error, even though it is assigned on the next line
	three = 3

	func(unused string) {
		fmt.Println("Unused arg. No compile error")
	}("what")
}
