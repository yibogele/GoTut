package main

import (
	"fmt"
	"time"
)

/**
 * author: will fan
 * created: 2019/5/3 6:31
 * description:
 */
func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom.")
			return
		default:
			fmt.Println("		.")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
