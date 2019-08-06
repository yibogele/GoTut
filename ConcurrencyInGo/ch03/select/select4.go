package main

import (
	"fmt"
	"time"
)

/**
 * created: 2019/7/30 11:04
 * By Will Fan
 */
func main() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
	loop:
		for {
			select {
			case <- done:
				break loop
			default:

			}

			workCounter++
			time.Sleep(1 * time.Second)
		}

	fmt.Printf("Achieved %v cycles of work before stop\n", workCounter)
}
