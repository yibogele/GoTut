package main

import (
	"fmt"
	"reflect"
	"strings"
)

/**
 * author: will fan
 * created: 2019/5/18 21:46
 * description:
 */
func main() {
	fruit := []string{
		"apple",
		"banana",
		"cherry",
	}
	fmt.Println(fruit)
	fmt.Println(reflect.TypeOf(fruit))
	fmt.Println(strings.TrimRight("rteABBA", "BA"))
	fmt.Println(strings.TrimLeft("etretABBA", "BA"))
}
