package main

import (
	"encoding/json"
	"fmt"
)

type MyData struct {
	One int
	two string
}

/**
 * created: 2019/5/8 9:41
 * By Will Fan
 */
func main() {
	in := MyData{1, "two"}
	fmt.Printf("%#v\n", in) // prints main.MyData{One:1, two:"two"}

	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded)) // prints {"One":1}

	var out MyData
	json.Unmarshal(encoded, &out)

	fmt.Printf("%#v\n", out) // prints main.MyData{One:1, two:""}
}
