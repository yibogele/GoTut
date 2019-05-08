package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
 * created: 2019/5/8 10:17
 * By Will Fan
 */
func main() {
	ch := make(chan int)
	done := make(chan struct{})
	for i := 0; i < 3; i++ {
		go func(idx int) {
			select {
			case ch <- (idx+1)*2: fmt.Println(idx, "sent result")
			case <- done: fmt.Println(idx, "exiting")
			}
		}(i)
		//runtime.Gosched()
		// GoSched will make different output
	}

	fmt.Println("result:", <-ch)
	close(done)
	time.Sleep(3 * time.Second)
}
