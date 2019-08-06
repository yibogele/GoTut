package main

import "fmt"

/**
 * created: 2019/7/30 13:13
 * By Will Fan
 */
func main() {
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data{
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData{
		fmt.Println(num)
	}
}
