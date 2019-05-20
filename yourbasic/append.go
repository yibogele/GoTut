package main

import "fmt"

/**
 * author: will fan
 * created: 2019/5/18 22:06
 * description:
 */
func main() {
	a := []byte("Hello: ")
	fmt.Println(len(a), cap(a))

	a1 := append(a, '1')
	a2 := append(a, '2')

	fmt.Println(string(a1), string(a2))
}
