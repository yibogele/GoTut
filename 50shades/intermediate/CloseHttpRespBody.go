package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
/**
 * created: 2019/5/8 10:43
 * By Will Fan
 */
func main() {
	resp, err := http.Get("https://api.ipify.org?format=json")
	defer resp.Body.Close()
	// not ok!, cause sometimes resp may be nil, and this will cause panic

	if err != nil{
		fmt.Println(err)
		return
	}

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
