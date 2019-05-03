package main

import "fmt"

/**
 * author: will fan
 * created: 2019/5/2 20:40
 * description:
 */
func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
