package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

/**
 * created: 2019/5/8 11:02
 * By Will Fan
 */
func main() {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	// ok, most of the time

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
