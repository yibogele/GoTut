package main

import "fmt"
import "time"

/**
 * created: 2019/5/8 10:07
 * By Will Fan
 */
func main() {
	ch := make(chan string)

	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
			time.Sleep(1 * time.Second)
		}
	}()

	ch <- "cmd.1"
	ch <- "cmd.2" // wont be processed
}
