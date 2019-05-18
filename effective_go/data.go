package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
	home string

}

func NewPerson(name string, age int, home string) *Person {
	p := new(Person)
	p.name = name
	p.age = age
	p.home = home
	return p
}

func NewPersonV2(name string, age int, home string) *Person {

	p := Person{name, age, home}
	return &p

}

func NewPersonV3(name string, age int, home string) *Person {
	return &Person{name, age, home}
}

const (
	Enone = iota
	Eio
	Einval
)
/**
 * created: 2019/5/18 13:33
 * By Will Fan
 */
func main() {
	a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}

	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(m)

	fmt.Println(NewPersonV3("",1,""))
	fmt.Println(NewPersonV3("", 1, ""))
}
