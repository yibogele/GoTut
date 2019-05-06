package main

import (
	"flag"
	"fmt"
)

/**
 * created: 2019/5/6 8:27
 * By Will Fan
 */

var name = flag.String("name", "World", "A name to say hello to.")
var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language.")
}

func main() {
	flag.Parse()

	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
