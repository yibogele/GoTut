package main

import "fmt"

/**
 * author: will fan
 * created: 2019/8/29 22:04
 * description:
 */
func main() {
	var array1 [3]*string

	array2 := [3]*string {new(string), new(string), new(string)}

	*array2[0] = "Red"
	*array2[1] = "Green"
	*array2[2] = "Blue"

	array1 = array2

	fmt.Println(array1)
}
