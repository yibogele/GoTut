package main

import (
	"fmt"
	"strings"
)

/**
 * created: 2019/5/8 14:55
 * By Will Fan
 */
func main() {
	str := "HI, I'M UPPER CASE!"

	lower := strings.ToLower(str)

	fmt.Println(lower)

	if strings.Contains(str, "case") {
		fmt.Println("Yes, exists!")
	}

	str = "abcdefghijklmnopqrstuvwxyz"
	fmt.Println("Char 3-10:", str[3:10])

	fmt.Println("First 5:", str[:5])

	fmt.Println("From 13 on:", str[13:])

	sentence := "I'm a sentence made up of words"
	words := strings.Split(sentence, " ")
	fmt.Printf("%v \n", words)

	fields := strings.Fields(sentence)
	fmt.Printf("%v \n", fields)

	str = "The blue whale loves blue fish."
	newstr := strings.Replace(str, "blue", "red", 1)
	fmt.Println(newstr)

	newstr = strings.ReplaceAll(str, "blue", "red")
	fmt.Println(newstr)

	path := "/home/fancx"
	isAbsolute := strings.HasPrefix(path, "/")
	isFancx := strings.HasSuffix(path, "fancx")
	fmt.Println(isAbsolute, isFancx)
}
