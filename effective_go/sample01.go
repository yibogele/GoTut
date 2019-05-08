package main

import (
	"fmt"
	"math/rand"
)


func main() {
	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}

	fmt.Println("My favorite number is ", rand.Intn(10))

	fmt.Println(a, s, m)
}

type T struct {
	name  string // name of the object
	value int    // its value
}
