package _chan

import "fmt"

/**
 * created: 2019/7/29 15:59
 * By Will Fan
 */
func main() {
	stringStream := make(chan string)

	go func() {
		stringStream <- "Hello channels"
	}()

	fmt.Println(<- stringStream)
}
