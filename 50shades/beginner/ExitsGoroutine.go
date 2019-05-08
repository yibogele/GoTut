package main

import (
	"fmt"
	"time"
)

/**
 * created: 2019/5/8 9:46
 * By Will Fan
 */
func main() {
	workerCount := 2

	for i := 0; i < workerCount; i++ {
		go doit(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("all done!")
}

func doit(workerId int) {
	fmt.Printf("[%v] is running\n", workerId)
	time.Sleep(3 * time.Second)
	fmt.Printf("[%v] is done\n", workerId)
}
