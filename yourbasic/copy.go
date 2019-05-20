package main

import "fmt"

/**
 * author: will fan
 * created: 2019/5/18 22:02
 * description:
 */
func main() {
	var dst, src []int
	src = []int{1,2,3}
	copy(dst, src)
	fmt.Println(dst)

	dst = append([]int(nil), src...)
	fmt.Println(dst)

	var dst2 []int
	dst2 = make([]int, len(src))
	copy(dst2, src)
	fmt.Println(dst2)
}
