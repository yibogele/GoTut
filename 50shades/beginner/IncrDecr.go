package beginner

import "fmt"

/**
 * created: 2019/5/8 9:25
 * By Will Fan
 */
func main() {
	data := []int {1,2,3}
	i := 0
	// wrong
	//++i
	//fmt.Println(data[i++])

	// right
	i++
	fmt.Println(data[i])
}
