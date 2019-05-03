package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

/**
 * author: will fan
 * created: 2019/5/2 20:32
 * description:
 */
func main() {
	go say("world")
	say("hello")
}
