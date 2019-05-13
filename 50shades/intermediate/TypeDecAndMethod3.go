package main

import (
	"sync"
)

type myLocker sync.Mutex

/**
 * created: 2019/5/13 9:46
 * By Will Fan
 */
func main() {
	var lock = new(sync.Mutex)
	lock.Lock()
	lock.Unlock()
}
