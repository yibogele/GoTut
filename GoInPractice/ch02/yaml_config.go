package main

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
)

/**
 * created: 2019/5/6 13:31
 * By Will Fan
 */
func main() {
	config, err := yaml.ReadFile("goinpractice/ch02/conf.yaml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))
}
