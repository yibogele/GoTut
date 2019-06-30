package main

import "fmt"

type Address struct {
	Street, City, State, Zip 	string
	IsShippingAddress			bool
}

type Customer struct {
	FirstName, LastName, Email, Phone	string
	Address								[]Address
}

func (c Customer) ToString() string {

}
/**
 * author: will fan
 * created: 2019/6/30 18:18
 * description:
 */
func main() {
	c := Customer{
		FirstName:	"Alex",
		LastName:	"John",
		Email:		"fcxlln@163.com",
		Phone:		"12345678",
		Address: 	[]Address{
			{
				Street:	"1 Mission Street",
				City: "San Francisco",
				State: "CA",
				Zip: "125328",
				IsShippingAddress: true,
			},
			{
				Street: "49 Stevenson Street",
				City: "San Francisco",
				State: "CA",
				Zip: "1231234",
			},
		},
	}

	fmt.Println(c)
}
