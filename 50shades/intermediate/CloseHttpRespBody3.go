package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

/**
 * created: 2019/5/8 11:06
 * By Will Fan
 */
func main() {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
