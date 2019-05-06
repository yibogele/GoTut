package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

/**
 * created: 2019/5/6 10:37
 * By Will Fan
 */
func main() {
	go echo(os.Stdin, os.Stdout)
	time.Sleep(30 * time.Second)
	fmt.Println("Timed out.")
	os.Exit(0)
}

func echo(in io.Reader, out io.Writer) {
	io.Copy(out, in)
}
