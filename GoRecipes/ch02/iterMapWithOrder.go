package main

import (
	"fmt"
	"sort"
)

/**
 * author: will fan
 * created: 2019/6/30 17:35
 * description:
 */
func main() {
	chapts := make(map[int]string)

	chapts[1] = "Beginning Go"
	chapts[2] = "Go Fundamentals"
	chapts[3] = "Structs and Interfaces"

	var keys []int
	for k := range chapts{
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, k := range keys{
		fmt.Println("Key:", k, "Value:", chapts[k])
	}
}
