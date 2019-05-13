package main

import (
	"fmt"
	"time"
)

/**
 * created: 2019/5/13 9:58
 * By Will Fan
 */
func main() {
	data := []string{"one", "two", "three"}

	for _, v := range data {
		vt := v
		go func() {
			fmt.Println(vt)
		}()
	}

	time.Sleep(3 * time.Second)
}
