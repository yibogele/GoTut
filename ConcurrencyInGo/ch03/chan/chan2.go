package _chan

import "fmt"

/**
 * created: 2019/7/29 16:07
 * By Will Fan
 */
func main() {
	intStream := make(chan int)
	close(intStream)

	integer, ok := <- intStream
	fmt.Printf("(%v) %v", ok, integer)
}
