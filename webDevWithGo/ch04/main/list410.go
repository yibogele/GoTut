package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
 * author: will fan
 * created: 2019/9/1 15:15
 * description:
 */
func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go web development")
}
func main() {

	http.HandleFunc("/welcome", messageHandler)

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
