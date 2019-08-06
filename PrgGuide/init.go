package main

import (
	"fmt"
	"time"
)

var a int

func init() {
	a = 1
}

var c = make(chan int, 10)
var b string

func f() {
	b = "hello, world"
	c <- 0
}
func Foo(n int) int {
	fmt.Println(n)
	return n
}

/**
 * created: 2019/7/15 9:26
 * By Will Fan
 */
func main() {
	//fmt.Println(a)

	go func() {
		fmt.Println(a)
		c <- 1
	}()

	go f()
	//fmt.Println(<- c)
	//fmt.Println(<- c)
	print(b)

	for i, ch := range "中国人" {
		fmt.Printf("%#U starts at byte position %d\n", ch, i)
	}

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}

	switch time.Now().Weekday() {
	case time.Sunday:
		fmt.Println("Today is Sunday")
	case time.Saturday:
		fmt.Println("Saturday")
	default:
		fmt.Println("Today is weekday")
	}

	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	c := 'a'
	switch c {
	case 'a', 'b', 'c':
		fmt.Println("3 basic")
	default:
		fmt.Println("default")
	}

	switch 2 {
	case 1:
		println("1")
		fallthrough
	case 2:
		println("2")
		fallthrough
	case 3:
		println("3")

	}

	switch Foo(1) {
	case Foo(0), Foo(1), Foo(2):
		fmt.Println("First case")
		fallthrough
	case Foo(3):
		fmt.Println("Second case")
	}
}
