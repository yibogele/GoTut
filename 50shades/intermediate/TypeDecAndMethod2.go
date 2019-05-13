package main

import (
	"sync"
)

type myLocker struct {
	sync.Mutex
}
/**
 * created: 2019/5/13 9:42
 * By Will Fan
 */
func main() {
	var lock myLocker
	lock.Lock()
	lock.Unlock()
}
