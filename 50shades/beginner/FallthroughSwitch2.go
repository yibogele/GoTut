package beginner

import "fmt"

/**
 * created: 2019/5/8 9:22
 * By Will Fan
 */
func main() {
	isSpace := func(ch byte)  bool {
		switch ch {
		case ' ', '\t':
			return true

		}
		return false
	}

	fmt.Println(isSpace('\t'))
	fmt.Println(isSpace(' '))
}
