package main

import "fmt"

type Suite int

const (
	Spades Suite = iota
	Hearts
	Diamonds
	Clubs
)

func (s Suite) String() string {
	return [...]string{"Spades", "Hearts", "Diamonds", "Clubs"}[s]
}

/**
 * created: 2019/7/15 13:01
 * By Will Fan
 */
func main() {
	s := Hearts
	fmt.Print(s)

	switch s {
	case Spades:
		fmt.Println(" are best")
	case Hearts:
		fmt.Println(" are second best")
	default:
		fmt.Println(" aren't very good")
	}
}
