package demo

import (
	"fmt"
	"sync"
)

/**
 * created: 2019/7/29 14:06
 * By Will Fan
 */
func main() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()

		fmt.Printf("Hello from %v!\n", id)
	}

	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i + 1)
	}
	wg.Wait()
}
