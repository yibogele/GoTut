package main

import "fmt"

func Swap(x, y int) (int, int) {
	return y, x
}

/**
 * author: will fan
 * created: 2019/6/30 13:48
 * description:
 */
func main() {
	x, y := Swap(2, 3)
	fmt.Println("Swap: ", x, y)
}
