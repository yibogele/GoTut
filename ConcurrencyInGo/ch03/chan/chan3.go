package _chan

import "fmt"

/**
 * created: 2019/7/29 16:11
 * By Will Fan
 */
func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
}
