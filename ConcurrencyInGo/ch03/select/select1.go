package main

/**
 * created: 2019/7/30 9:00
 * By Will Fan
 */
func main() {
	var c1, c2 <-chan interface{}
	var c3 chan<- interface{}

	select {
	case <- c1:
		//
	case <- c2:
		//
	case c3 <- struct {}{}:
		//
	}
}
