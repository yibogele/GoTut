package main

import "fmt"

/**
 * created: 2019/7/15 12:38
 * By Will Fan
 */
func main() {
	s := foo()
	fmt.Println(s)
	b := bar()
	fmt.Println(b)
}

func foo() string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recovering from: %q\n", err)
		}
	}()

	s := "before "
	panic("disaster")
	s += "after"
	return s
}

func bar() (s string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recovering from %q\n", err)
			s += "during"
		}
	}()

	s = "before "
	panic("disaster")
	s += "after"
	return s
}