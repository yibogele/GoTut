package main

import "fmt"


/**
 * author: will fan
 * created: 2019/6/30 14:22
 * description:
 */
func main() {
	var x [5]int

	x[0] = 1
	x[4] = 25
	fmt.Println("X:", x)

	x[1] = 10
	x[2] = 23
	x[3] = 13
	fmt.Println("X: ", x)

	y := [5]int{1,2,3,4,5}
	fmt.Println("Y: ", y)

	z := [...]int{6,7,8,9,10}
	fmt.Println("Z: ", z)
	fmt.Println("Length of z: ", len(z))

	langs := [4]string{0:"Go", 3:"Python"}
	fmt.Println("Value of langs: ", langs)
	langs[1] = "Scala"
	langs[2] = "Java"

	for _, lang := range langs{
		fmt.Println(lang)
	}

}
