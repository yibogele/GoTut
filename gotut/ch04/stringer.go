package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
/**
 * author: will fan
 * created: 2019/5/2 20:10
 * description:
 */
func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaddsdsf Broc", 1001}
	fmt.Println(a, z)
}
