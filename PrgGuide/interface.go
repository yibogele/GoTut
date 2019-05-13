package main

import (
	"fmt"
	"strconv"
)

type MyStringer interface {
	String() string
}

type Temp int

func (t Temp) String() string {
	return strconv.Itoa(int(t)) + " ℃"
}

type Point struct {
	x, y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}
/**
 * created: 2019/5/13 16:08
 * By Will Fan
 */
func main() {
	var x MyStringer

	x = Temp(24)
	fmt.Println(x.String())

	x = &Point{1,2}
	fmt.Println(x.String())

	var y interface{}

	y = 2.4
	fmt.Println(y)

	y = &Point{1,2}
	fmt.Println(y)

}
