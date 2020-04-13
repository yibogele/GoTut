package main

import (
	"flag"
	"fmt"
)

/**
 * author: will fan
 * created: 2020/1/14 20:57
 * description:
 */
var name = flag.String("name", "World", "A name to say hello to.")

var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use spanish language.")
}

func main() {
	flag.Parse()

	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)

	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
