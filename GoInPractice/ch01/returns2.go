package main

import "fmt"
/**
 * created: 2019/5/5 16:45
 * By Will Fan
 */

func Names() (first, second string) {
	first = "Foo"
	second = "Bar"
	return
}
func main() {
	n1, n2 := Names()
	fmt.Println(n1, n2)
}
