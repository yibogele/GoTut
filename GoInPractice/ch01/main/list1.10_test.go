package main

import "testing"

/**
 * author: will fan
 * created: 2020/1/13 21:04
 * description:
 */
func TestNames(t *testing.T) {
	name := getName()

	if name != "Hello" {
		t.Error("Response from getName is unexpected value")
	}
}
