package main

import (
	"compress/gzip"
	"io"
	"os"
)

/**
 * created: 2019/5/6 14:37
 * By Will Fan
 */
func main() {
	for _, file := range os.Args[1:] {
		_ = compress(file)
	}
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