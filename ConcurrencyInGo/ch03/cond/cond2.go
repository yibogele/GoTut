package cond

import (
	"fmt"
	"sync"
)

/**
 * created: 2019/7/29 14:48
 * By Will Fan
 */
func main() {
	type Button struct {
		Clicked * sync.Cond
	}

	button := Button{ Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing Window.")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast()

	clickRegistered.Wait()

}
