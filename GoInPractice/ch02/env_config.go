package main

import (
	"fmt"
	"net/http"
	"os"
)

/**
 * created: 2019/5/6 14:03
 * By Will Fan
 */
func main() {
	http.HandleFunc("/", homePage)
	_ = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	_, _ = fmt.Fprint(res, "The homepage.")
}
