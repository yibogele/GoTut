package main

import (
	"fmt"
	"time"
)

/**
 * created: 2019/7/30 15:38
 * By Will Fan
 */
func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(chanels ...<-chan interface{}) <-chan interface{} {
		switch len(chanels) {
		case 0:
			return nil
		case 1:
			return chanels[0]
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)

			switch len(chanels) {
			case 2:
				select {
				case <-chanels[0]:
				case <-chanels[1]:

				}
			default:
				select {
				case <-chanels[0]:
				case <-chanels[1]:
				case <-chanels[2]:
				case <-or(append(chanels[3:], orDone)...):
				}
			}
		}()

		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)

			time.Sleep(after)
		}()

		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}
