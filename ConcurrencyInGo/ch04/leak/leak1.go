package main

import "fmt"

/**
 * created: 2019/7/30 15:11
 * By Will Fan
 */
func main() {
	doWork := func(strings <-chan string) <-chan interface {}{
		completed := make(chan interface{})

		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)

			for s := range strings{
				fmt.Println(s)
			}
		}()

		return completed
	}

	doWork(nil)
	fmt.Println("Done.")
}
