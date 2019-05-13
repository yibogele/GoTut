package main

import (
	"fmt"
)

/**
 * created: 2019/5/13 9:50
 * By Will Fan
 */
func main() {
	loop:
		for {
			switch {
			case true:
				fmt.Println("breaking out...")
				break loop
			}
		}

	fmt.Println("out!")
}
