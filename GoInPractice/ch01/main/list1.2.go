package main


/**
 * author: will fan
 * created: 2020/1/12 21:00
 * description:
 */
func NamesV2() (first, second string) {
	first = "Foo"
	second = "Bar"
	return
}

func main() {
	n1, n2 := NamesV2()
	println(n1, n2)
}
