package main

import "fmt"

/**
 * author: will fan
 * created: 2019/8/29 21:57
 * description:
 */
func main() {
	var array1 [5]string
	array2 := [5]string{"Red", "Blue", "Green", "Yellow"}

	array1 = array2

	fmt.Println(array1)
	fmt.Println(array2)
}
