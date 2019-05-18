package main

import "fmt"

/**
 * author: will fan
 * created: 2019/5/8 20:29
 * description:
 */
func main() {
	buf := []rune("hello")
	buf[0] = 'H'
	s := string(buf)
	fmt.Println(s)

	a := []byte("ba")

	a1 := append(a, 'd')
	a2 := append(a, 'g')

	fmt.Println(string(a1)) // bag
	fmt.Println(string(a2)) // bag

	const name = 9876543210 * 9876543210
	fmt.Println(float64(name))
}
