package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
 * author: will fan
 * created: 2019/9/1 15:00
 * description:
 */
func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to GO Web Development")
}
func main() {
	mux := http.NewServeMux()

	mh := http.HandlerFunc(messageHandler)
	mux.Handle("/welcome", mh)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
