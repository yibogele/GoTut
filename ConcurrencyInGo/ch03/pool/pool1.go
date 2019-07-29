package pool

import (
	"fmt"
	"sync"
)

/**
 * created: 2019/7/29 15:25
 * By Will Fan
 */
func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct {}{}
		},
	}

	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()
}
