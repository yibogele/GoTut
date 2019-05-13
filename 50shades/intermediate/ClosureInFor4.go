package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

/**
 * created: 2019/5/13 10:08
 * By Will Fan
 */
func main() {
	//data := []field{{"one"}, {"two"}, {"three"}}
	data := []*field{{"one"}, {"two"}, {"three"}}

	for _,v := range data {
		//v := v
		go v.print()
	}

	time.Sleep(3 * time.Second)
}
