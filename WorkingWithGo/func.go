package main

import (
	"fmt"
	"math"
)

/**
 * created: 2019/5/8 15:29
 * By Will Fan
 */
func main() {
	sum1 := sum(1,3,4,5)
	fmt.Println(sum1)

	eles := []int{1,2,3,4}
	fmt.Println(sum(eles...))
}

func echo(str string) {
	fmt.Println(str)
}

func say(s string) string {
	return "Hello" + s
}

func say2(s string) (r string) {
	r = "hello" + s
	return
}

func divide(x, y float64) (float64, float64) {
	q := math.Trunc(x/y)
	r := math.Mod(x, y)
	return q, r
}

func divide2(x, y float64) (q, r float64) {
	q = math.Trunc(x/y)
	r = math.Mod(x, y)
	return
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}