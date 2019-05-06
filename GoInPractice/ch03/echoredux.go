package main

import (
	"fmt"
	"os"
	"time"
)

/**
 * created: 2019/5/6 16:30
 * By Will Fan
 */
func main() {
	done := time.After(30 * time.Second)
	echo := make(chan []byte)
	go readStdin(echo)
	for {
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("Timed out.")
			os.Exit(0)
		}
	}
}

func readStdin(out chan<- []byte) {
	for {
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			out <- data
		}
	}
}
