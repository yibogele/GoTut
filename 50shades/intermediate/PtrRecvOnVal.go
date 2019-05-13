package main

import (
	"fmt"
)

type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}
/**
 * created: 2019/5/13 12:55
 * By Will Fan
 */
func main() {
	d1 := data{"one"}
	d1.print()

	var in printer = data{"two"}
	in.print()

	m := map[string]data{"x": {"three"}}
	m["x"].print()

	s := []string{"f", "o"}
	a := [...]string{"f", "o"}
	n := map[string]int {
		"m":1,
		"n":2,
	}
	_ = s
	_ = a
	_ = n
}
