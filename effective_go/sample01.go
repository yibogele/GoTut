package main

import (
	"fmt"
	"math/rand"
)
func main() {
	fmt.Println("My favorite number is ", rand.Intn(10))
}

type T struct {
	name  string // name of the object
	value int    // its value
}
