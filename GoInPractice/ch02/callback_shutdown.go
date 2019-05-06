package main

import (
	"fmt"
	"net/http"
	"os"
)

/**
 * created: 2019/5/6 14:10
 * By Will Fan
 */
func main() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/", homtpage)
	_ = http.ListenAndServe(":8080", nil)
}

func shutdown(res http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}

func homtpage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	_, _ = fmt.Fprint(res, "the homepage.")
}
