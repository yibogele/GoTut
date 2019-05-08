package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("Received %d\n", i)
	}
	wg.Done()
}
/**
 * created: 2019/5/7 9:56
 * By Will Fan
 */
func main() {
	c := make(chan int)
	go printer(c)
	wg.Add(1)

	for i := 1; i < 10; i++ {
		c <- i
	}

	close(c)
	wg.Wait()
}
