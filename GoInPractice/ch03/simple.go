package main

import (
	"fmt"
	"runtime"
)

/**
 * created: 2019/5/6 10:46
 * By Will Fan
 */
func main() {
	fmt.Println("Outside a goroutine.")
	go func() {
		fmt.Println("Inside a goroutine.")
	}()
	fmt.Println("Outside again.")

	runtime.Gosched()
}
