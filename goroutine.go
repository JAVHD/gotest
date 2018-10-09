package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name  string
	sex   int
	skill string
}

func (p Person) Say() {
	//return "name:"+p.name+",age:"+strconv.Itoa(p.sex)
	fmt.Println("name:" + p.name + ",age:" + strconv.Itoa(p.sex))
}

func (p Person) Do() {
	fmt.Println("My skill is " + p.skill)
}

func main() {

	bear := Person{"bearxxx", 18, "golangxxx"}

	//bear.Say()
	go bear.Say() //go关键字开辟一个新的goroutine
	bear.Do()

}
