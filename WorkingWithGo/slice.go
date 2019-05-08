package main

import "fmt"

/**
 * created: 2019/5/8 15:10
 * By Will Fan
 */
func main() {
	var empty []int
	_ = empty

	alphas := []string{"one", "two", "three", "four"}
	fmt.Println(alphas)

	alphas = append(alphas, "five", "six")
	fmt.Println(alphas)

	fmt.Println(existsElements("five", alphas))
}

func existsElements(ele string, eles []string) bool  {

	for _, v := range eles {
		if v == ele {
			return true
		}
	}
	return false

}
