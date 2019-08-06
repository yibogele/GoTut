package main

import "fmt"



/**
 * created: 2019/7/15 13:17
 * By Will Fan
 */
func main() {
	m := make(map[string]float64)
	m["pi"] = 3.14
	fmt.Println(m["pi"])

	v1 := m["pi"]
	v2 := m["foo"]
	fmt.Println(v1, v2)

	//_, exists := m["pi"]
	if x, ok := m["pi"]; ok {
		fmt.Println(x)
	}

	delete(m, "pi")
	fmt.Println(m)
}
