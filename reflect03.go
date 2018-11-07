/*
* @Author: bear
* @Date:   2018-10-27 22:21:47
* @Last Modified by:   bear
* @Last Modified time: 2018-11-06 16:42:53
 */
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string
	age  int
}

type Manager struct {
	User
	title string
}

func main() {

	var a = User{"bear", 18}
	var b = Manager{User{"zhang", 25}, "hello"}

	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(a).Kind(), reflect.TypeOf(a).Name())

	fmt.Println(reflect.TypeOf(&a), reflect.TypeOf(&b))
	fmt.Println(reflect.TypeOf(&a).Kind(), reflect.TypeOf(&b).Kind())

	//reflect.

	//int ptr

	//reflect.
	fmt.Println(reflect.TypeOf(&a))
	fmt.Println(reflect.TypeOf(&a).Elem())

	// t := reflect.TypeOf(a)
	// t.FieldByIndex()

	//匿名字段

}
