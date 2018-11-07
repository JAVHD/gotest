/*
* @Author: bear
* @Date:   2018-11-07 09:13:05
* @Last Modified by:   bear
* @Last Modified time: 2018-11-07 09:13:07
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {

	//var age int
	//
	//age = 18
	//
	//fmt.Println(reflect.TypeOf(age).Kind())
	//
	//fmt.Println(reflect.TypeOf(&age).Kind())
	//
	//fmt.Println(reflect.TypeOf(&age).Elem())
	//
	//fmt.Println(reflect.TypeOf(age).Elem())

	type user struct {
		name string
		age  int
	}

	u := user{"age", 18}

	tou := reflect.TypeOf(u)

	fmt.Println(reflect.TypeOf(u))

	fmt.Println(reflect.TypeOf(u).Kind())

	fmt.Println(reflect.TypeOf(&u).Kind())
	//
	fmt.Println(reflect.TypeOf(&u).Elem())
	//
	//fmt.Println(reflect.TypeOf(u).Elem())

	//a, _ := fmt.P

	for i := 0; i < tou.NumField(); i++ {
		fmt.Println(tou.Field(i).Name)
	}

}
