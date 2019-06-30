package main

import "fmt"

/**
 * author: will fan
 * created: 2019/6/30 17:24
 * description:
 */
func main() {
	var chapts map[int]string

	chapts = make(map[int]string)

	chapts[1] = "Beginning Go"
	chapts[2] = "Go Fundamentals"
	chapts[3] = "Structs and Interfaces"

	for k,v := range chapts{
		fmt.Printf("Key: %d Value: %s\n", k, v)
	}

	langs := map[string]string{
		"EL": "Greek",
		"EN": "English",
		"ES": "Spanish",
		"FR": "French",
		"HI": "Hindi",
	}
	delete(langs, "EL")

	if lang, ok := langs["EL"]; ok {
		fmt.Println(lang)
	}else {
		fmt.Println("Key don't exist")
	}
}
