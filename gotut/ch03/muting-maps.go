package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["answer"] = 12
	fmt.Println("The value: ", m["answer"])

	m["answer"] = 24
	fmt.Println("The value: ", m["answer"])

	delete(m, "answer")
	fmt.Println("The value: ", m["answer"])

	v, ok := m["answer"]
	fmt.Println("The value:", v, "present?", ok)

}
