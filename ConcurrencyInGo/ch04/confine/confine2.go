package main

import "fmt"

/**
 * created: 2019/7/30 13:17
 * By Will Fan
 */
func main() {
	chanOwner := func() <-chan int{
		result := make(chan int, 5)

		go func() {
			defer close(result)
			for i := 0; i<=5; i++ {
				result <- i
			}
		}()

		return result
	}

	consumer := func(result <-chan int) {
		for res := range result{
			fmt.Printf("Received %v\n", res)
		}

		fmt.Println("Done received")
	}

	results := chanOwner()
	consumer(results)



}



