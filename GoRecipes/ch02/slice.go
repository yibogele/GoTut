package main

import "fmt"


/**
 * author: will fan
 * created: 2019/6/30 17:05
 * description:
 */
func main() {
	x := []int{10, 20, 30}
	fmt.Println("[Slice:x] length is %d capacity is %d\n", len(x), cap(x))

	y := make([]int, 5, 10)
	copy(y, x)
	fmt.Println("[Slice:y] length is %d capacity is %d\n", len(y), cap(y))
	fmt.Println("[Slice:y] is ", y)
	y[3] = 40
	y[4] = 50
	fmt.Println("[Slice y:", y)
}
