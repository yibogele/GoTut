package main

import "fmt"

func fib(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
/**
 * author: will fan
 * created: 2019/5/2 20:42
 * description:
 */
func main() {
	c := make(chan int, 10)
	go fib(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
