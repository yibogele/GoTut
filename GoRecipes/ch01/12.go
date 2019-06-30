package main

import "fmt"

const Title = "Person Details"

var Country = "USA"


func main() {
	fname, lname := "Changxue", "Fan"
	age := 35

	fmt.Println(Title)

	fmt.Println("First Name: ", fname)
	fmt.Println("Last Name: ", lname)
	fmt.Println("Age: ", age)
	fmt.Println("Country: ", Country)
}
