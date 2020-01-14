package main

import "fmt"

/**
 * author: will fan
 * created: 2020/1/13 21:02
 * description:
 */
func getName() string {
	return "Hello"
}
func main() {
	name := getName()
	fmt.Println(name)
}
