package main

import "fmt"

/**
 * author: will fan
 * created: 2020/4/14 11:09
 * description:
 */

func main() {

	f := func() {
		i := 0
		fmt.Println(i)
	}

	for  {
		f()
		if true{
			break
		}
	}
}
