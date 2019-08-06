package main

import (
	"runtime"
	"sync"
	"time"
)

func init() {
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)
}
/**
 * created: 2019/7/15 16:42
 * By Will Fan
 */
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		wg.Done()
	}()

	go func() {
		wg.Done()
	}()

	wg.Wait()

	ch := make(chan struct{})
	select {
	case <-ch:
		println("ch received")
	case <- time.After(5 * time.Second):
		println(" time out")
	}
}
