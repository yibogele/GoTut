package main

import (
	"log"
	"os"

	_ "./matchers"
	"./search"
)

func init() {
	log.SetOutput(os.Stdout)
}
/**
 * created: 2019/5/7 11:40
 * By Will Fan
 */
func main() {
	search.Run("president")
}
