package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/**
 * author: will fan
 * created: 2019/9/1 15:32
 * description:
 */
func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there")
}
func main() {
	http.HandleFunc("/welcome", messageHandler)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
	}

	log.Println("Listening...")
	server.ListenAndServe()
}
