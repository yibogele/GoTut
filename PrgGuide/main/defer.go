package main

import "fmt"

/**
 * author: will fan
 * created: 2020/4/14 11:55
 * description:
 */
func main() {
	defer fmt.Println("World1")
	fmt.Println("Hello1")

	// panic
	//panicDefer := func() {
	//	defer fmt.Println("World2")
	//	panic("Eeek")
	//	fmt.Println("Hello2")
	//}
	//panicDefer()

	// defer order
	deferOrder := func() {
		fmt.Println("start")
		for i:= 1; i<=3; i++ {
			defer fmt.Println(i)
		}
		fmt.Println("end")
	}
	deferOrder()

	// change return value of named param
	Foo := func() (result int) {
		defer func() {
			result = 7
		}()

		return 0
	}
	fmt.Println(Foo())
}
