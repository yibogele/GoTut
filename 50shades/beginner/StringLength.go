package beginner

import (
	"fmt"
	"unicode/utf8"
)

/**
 * created: 2019/5/8 8:45
 * By Will Fan
 */
func main() {
	data := "♤"
	fmt.Println(len(data))
	fmt.Println(utf8.RuneCountInString(data))
}
