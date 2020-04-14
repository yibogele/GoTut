package main

import (
	"fmt"
	"time"
)

/**
 * author: will fan
 * created: 2020/4/14 11:15
 * description:
 */
func main() {
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Today is Saturday")
	case time.Sunday:
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Today is weekday")
	}

	//
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	//case list
	whiteSpace := func(c rune) bool{
		switch c {
		case ' ', '\t', '\n', '\f', '\r':
			return true
		}
		return false
	}
	fmt.Println(whiteSpace(' '))

	// fall through
	switch 2 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}

	// execution order
	Foo := func(n int) int {
		fmt.Println(n)
		return n
	}
	fmt.Println("////////////////")
	switch Foo(1) {
	case Foo(0), Foo(1), Foo(2):
		fmt.Println("First case")
		fallthrough
	case Foo(3):
		fmt.Println("Second case ")
	}
}
