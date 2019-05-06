package main

import "time"

/**
 * created: 2019/5/6 16:55
 * By Will Fan
 */
func main() {
	ch := make(chan bool)
	timeout := time.After(600 * time.Millisecond)
	go send(ch)
	for {
		select {
		case <-ch:
			println("Got message")
		case <-timeout:
			println("Timed out")
			return
		default:
			println("yawn")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func send(ch chan bool) {
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch)
	println("Send and close")
}