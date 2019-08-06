package main

import "fmt"

type Person struct {
	name string
	age int
}

/**
 * created: 2019/7/15 13:06
 * By Will Fan
 */
func main() {
	var x Person
	x.age = 12

	var p *Person
	p = new(Person)
	p.age = 12

	y := Person{
		"bob",
		13,
	}

	jim := Person{
		name: "Jim",
	}

	pjohn := &Person{
		name: "John",
	}

	fmt.Println(y, jim, pjohn)
}
