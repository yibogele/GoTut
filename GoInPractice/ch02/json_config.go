package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	Enabled	bool
	Path 	string
}
/**
 * created: 2019/5/6 11:43
 * By Will Fan
 */
func main() {
	file, _ := os.Open("goinpractice/ch02/config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error: ", err)

	}
	fmt.Println(conf.Path)
}
