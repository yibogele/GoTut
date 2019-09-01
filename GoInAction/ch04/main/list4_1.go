package main

import "fmt"

/**
 * author: will fan
 * created: 2019/8/28 21:52
 * description:
 */
func main() {
	var array [5]int
	//_ = array

	array = [5]int {1, 2, 3, 4, 5}
	array = [...]int {1, 2, 3, 4, 5}
	array = [5]int {1: 2, 3: 4}

	fmt.Println(array)

	array[3] = 123
	fmt.Println(array)

	parray := [5]*int {0: new(int), 1: new(int)}
	fmt.Println(parray)

	*parray[0] = 0
	*parray[1] = 10

}
