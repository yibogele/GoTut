package main

import "fmt"

/**
 * created: 2019/5/16 9:38
 * By Will Fan
 */
func main() {
	//var m map[string]float64
	m := make(map[string]float64, 100)
	m["pi"] = 3.14
	fmt.Println(m["pi"])
}
