package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

func (u User) MyHander() string {

	return fmt.Sprintf("%+v", u)
}

func main() {

	//m := User{"bear", 18}

	var m User

	m.name = "bear"
	m.age = 18

	fmt.Println(m.MyHander())

}
