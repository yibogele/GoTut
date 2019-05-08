package main

import (
	"fmt"
	"time"
)

/**
 * created: 2019/5/8 10:24
 * By Will Fan
 */
func main() {
	var ch chan int
	for i:= 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1)*2
		}(i)
	}

	//
	fmt.Println("result:", <-ch)
	time.Sleep(2*time.Second)
}
