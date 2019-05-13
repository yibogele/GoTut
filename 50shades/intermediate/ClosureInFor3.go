package main

import (
	"fmt"
	"time"
)

/**
 * created: 2019/5/13 10:01
 * By Will Fan
 */
func main() {
	data := []string{"one", "two", "three"}

	for _, v := range data {
		go func( val string) {
			fmt.Println(val)
		}(v)
	}

	time.Sleep(3 * time.Second)
}
