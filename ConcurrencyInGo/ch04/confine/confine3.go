package main

import (
	"bytes"
	"fmt"
	"sync"
)

/**
 * created: 2019/7/30 13:32
 * By Will Fan
 */
func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data{
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[3:])
	go printData(&wg, data[:3])

	wg.Wait()
}
