package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a int = 123
	var b float32 = 1.23
	var c string = "hello"
	var d []string = []string{"my", "name", "is", "bear"}

	type Foo struct {
		X string
		Y int
	}

	fmt.Println(reflect.TypeOf(a)) //int
	fmt.Println(reflect.TypeOf(b)) //float32
	fmt.Println(reflect.TypeOf(c)) //string
	fmt.Println(reflect.TypeOf(d)) //[]string

	var foo Foo

	fmt.Println(reflect.TypeOf(foo)) //main.Foo
	fmt.Println(reflect.ValueOf(a))

}
