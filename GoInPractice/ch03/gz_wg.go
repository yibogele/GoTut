package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

/**
 * created: 2019/5/6 14:48
 * By Will Fan
 */
func main() {
	var wg sync.WaitGroup
	var i int = -1
	var file string

	for i, file = range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			compress(filename)
			wg.Done()
		}(file)
	}

	wg.Wait()
	fmt.Printf("compressed %n files", i + 1)
}

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	_ = gzout.Close()

	return err
}

