package main

import (
	"github.com/NasSilverBullet/golang-interface-practice/pkg"
)

func main() {
	u := pkg.NewUser()
	(*u).SayHi()
	(*u).SayBye()
}
