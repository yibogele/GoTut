package main

import "fmt"

/**
 * author: will fan
 * created: 2019/9/14 16:57
 * description:
 */
func main() {
	list := user{
		name:       "List",
		email:      "list@163.com",
		ext:        123,
		privileged: false,
	}
	_ = list

	john := user{"John", "john@126.com", 123, true}
	_ = john

	fan := admin{
		person: user{
			name:       "fan",
			email:      "fan@126.com",
			ext:        0,
			privileged: false,
		},
		level:  "super",
	}
	_ = fan

	var dur	Duration
	dur = 1000
	_ = dur

	bill := user{"Bill", "bill@123.com", 123, true}
	bill.notify()

	kate := &user{"Kate", "kate@1234.com", 123, false}
	kate.changeEmail("kate@qq.com")

	kate.notify()
}

type user struct {
	name		string
	email 		string
	ext 		int
	privileged 	bool
}

var bill user

type admin struct {
	person	user
	level	string
}

type Duration int64

func (u user) notify()  {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email

}