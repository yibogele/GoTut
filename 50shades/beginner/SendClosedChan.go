package main

import (
	"fmt"
	"time"
)

/**
 * created: 2019/5/8 10:12
 * By Will Fan
 */
func main() {
	ch := make(chan int)
	for i := 0; i <3; i++ {
		go func(idx int) {
			ch <- (idx + 1)*2
		}(i)
	}

	fmt.Println(<-ch)
	close(ch)

	time.Sleep(2 * time.Second)
}
