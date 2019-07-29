package main

import (
	"fmt"
	"sync"
)

/**
 * created: 2019/7/29 16:16
 * By Will Fan
 */
func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			<- begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("unblocking goroutines...")
	close(begin)
	wg.Wait()
}
