package main

import "fmt"

/**
 * author: will fan
 * created: 2019/6/30 17:10
 * description:
 */
func main() {
	x := make([]int, 2, 5)
	x[0] = 10
	x[1] = 20
	fmt.Println("Slice x: ", x)

	x = append(x, 30, 40, 50)
	fmt.Println("Slice x:", x)

	x = append(x, 60, 70)
	fmt.Println("Slice x:", x)

}
