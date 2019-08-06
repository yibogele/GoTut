package main

import (
	"fmt"
	"net/http"
)

/**
 * created: 2019/7/30 16:15
 * By Will Fan
 */
func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan *http.Response {
		response := make(chan *http.Response)

		go func() {
			defer close(response)

			for _, url := range urls{
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					continue
				}

				select {
				case <-done:
					return
				case response <- resp:
				}
			}
		}()

		return response
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"www.baidu.com", "https://badhost"}
	for response := range checkStatus(done, urls...){
		fmt.Printf("respones %v\n", response)
	}
}
