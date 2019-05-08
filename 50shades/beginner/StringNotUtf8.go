package beginner

import (
	"fmt"
	"unicode/utf8"
)

/**
 * created: 2019/5/8 8:34
 * By Will Fan
 */
func main() {
	data1 := "ABC"
	fmt.Println(utf8.ValidString(data1))

	data2 := "A\xfeC"
	fmt.Println(utf8.ValidString(data2))
}
