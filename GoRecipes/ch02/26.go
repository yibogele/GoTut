package main

import "fmt"

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums{
		total += num
	}
	return total
}
/**
 * author: will fan
 * created: 2019/6/30 13:51
 * description:
 */
func main() {
	total := Sum(1,2,3,4,5)
	fmt.Println("Total: ", total)


}
