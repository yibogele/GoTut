package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(), "it didnt work.",
	}
}
/**
 * author: will fan
 * created: 2019/5/2 20:13
 * description:
 */
func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
