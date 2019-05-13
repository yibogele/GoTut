package main

import (
	"fmt"
	"time"
)

/**
 * created: 2019/5/13 9:54
 * By Will Fan
 */
func main() {
	data := []string{"One", "two", "three"}

	for _, v := range data {
		go func() {
			fmt.Println(v)
		}()
	}

	time.Sleep(3 * time.Second)
}
