package main

import "fmt"

/**
 * author: will fan
 * created: 2020/4/14 11:01
 * description:
 */
func main() {
	a := []string{"Foo", "Bar"}
	for i, s := range a{
		fmt.Println(i, s)
	}

	//
	for i, ch := range "中国香港"{
		fmt.Printf("%#U starts at byte point %d\n", ch, i)
	}

	// map
	m := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}
	for k, v := range m{
		fmt.Println(k, v)
	}

	// channel
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	for n := range ch{
		fmt.Println(n)
	}
}
