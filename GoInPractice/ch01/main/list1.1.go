package main

import "fmt"
/**
 * author: will fan
 * created: 2020/1/12 20:50
 * description:
 */
func Names() (string, string) {
	return "Foo", "Bar"
}
func main() {
	n1, n2 := Names()
	fmt.Println(n1, n2)

	n3, _ := Names()
	fmt.Println(n3)
}
