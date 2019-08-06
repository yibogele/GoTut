package main

/**
 * created: 2019/7/15 15:43
 * By Will Fan
 */
func main() {
	ch := make(chan int)
	ch <- 1
	println(<-ch)
}
