package main

import (
	"fmt"
	"gopkg.in/gcfg.v1"
)

/**
 * created: 2019/5/6 13:56
 * By Will Fan
 */
func main() {
	config := struct {
		Section struct{
			Enabled	bool
			Path	string
		}
	}{}

	err := gcfg.ReadFileInto(&config, "goinpractice/ch02/conf.ini")
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}

	fmt.Println(config.Section.Path, config.Section.Enabled)
}
