package demo

import (
	"fmt"
	"sync"
)

/**
 * created: 2019/7/29 12:58
 * By Will Fan
 */
func main() {
	var wg sync.WaitGroup

	for _, salution := range []string{"hello", "greetings", "good day"}{
		wg.Add(1)
		go func(sal string) {
			defer wg.Done()
			fmt.Println(sal)
		}(salution)
	}
	wg.Wait()
}
