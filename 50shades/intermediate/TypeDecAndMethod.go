package main

import (
	"sync"
)

type MyMutex sync.Mutex

/**
 * created: 2019/5/13 9:38
 * By Will Fan
 */
func main() {
	var mtx MyMutex

	mtx.Lock()
	mtx.Unlock()
}
