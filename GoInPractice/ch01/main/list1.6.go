package main

import (
	"fmt"
	"time"
)

/**
 * author: will fan
 * created: 2020/1/13 20:47
 * description:
 */
func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	go count()

	time.Sleep(time.Millisecond * 2)
	fmt.Println("Hello World")
	time.Sleep(time.Millisecond * 5)
}
