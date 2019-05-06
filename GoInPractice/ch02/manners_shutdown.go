package main

import (
	"fmt"
	"github.com/braintree/manners"
	"net/http"
	"os"
	"os/signal"
)

/**
 * created: 2019/5/6 14:16
 * By Will Fan
 */
func main() {
	handler := newHandler()
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutdown(ch)

	_ = manners.ListenAndServe(":8080", handler)
}

type handler struct {

}

func newHandler() *handler {
	return &handler{}
}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request)  {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "fancx"
	}

	fmt.Fprint(res, "hello, my name is ", name)
	
}

func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}