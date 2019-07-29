package main

import (
	"fmt"
	"sync"
)

/**
 * created: 2019/7/29 12:55
 * By Will Fan
 */
func main() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()

		salutation = "welcome"

	}()

	wg.Wait()

	fmt.Println(salutation)
}
