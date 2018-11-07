/*
* @Author: bear
* @Date:   2018-11-05 22:57:12
* @Last Modified by:   bear
* @Last Modified time: 2018-11-06 10:45:13
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {

	type user struct {
		name string
		age  int
	}

	type manage struct {
		user
		title string
	}

	var u user
	var m manage

	ut := reflect.TypeOf(u)
	mt := reflect.TypeOf(m)

	fmt.Println(ut, mt)

	mtName, _ := mt.FieldByName("user")

	fmt.Println(mtName)
	fmt.Println(mtName.Name, mtName.Type)

	for i := 0; i < mt.NumField(); i++ {
		fmt.Println(mt.Field(i).Name)
	}

}
