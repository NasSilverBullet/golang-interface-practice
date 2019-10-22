package pkg

import "fmt"

var u *User

func init() {
	var utmp User
	utmp = &user{}
	u = &utmp
}

type User interface {
	SayHi()
	SayBye()
}

func NewUser() *User {
	return u
}

type user struct {
}

func (u *user) SayHi() {
	fmt.Println("Hi")
}

func (u *user) SayBye() {
	fmt.Println("Bye")
}
