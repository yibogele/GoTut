package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

/**
 * author: will fan
 * created: 2019/5/2 20:05
 * description:
 */
func main() {
	var w Writer

	w = os.Stdout

	fmt.Fprint(w, "hello, writer\n")
}
