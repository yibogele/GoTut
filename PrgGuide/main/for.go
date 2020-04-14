package main

import "fmt"

/**
 * author: will fan
 * created: 2020/4/14 10:54
 * description:
 */
func main() {
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}

	fmt.Printf("%d\n", sum)

	/////while
	power := 1
	for power < 5 {
		power *= 2
	}
	fmt.Println(power)

	// infinite loop
	//sum = 0
	//for {
	//	sum++
	//}
	//fmt.Println(sum)

	// foreach
	strings := []string{"hello", "world"}
	for i, s := range strings{
		fmt.Println(i, s)
	}

	// break, continue
	sum = 0
	for i := 0; i < 5; i++ {
		if i % 2 != 0 {
			continue
		}
		sum += i
	}
	fmt.Println(sum)
}
