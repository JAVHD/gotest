/*
* @Author: bear
* @Date:   2018-11-07 11:13:01
* @Last Modified by:   bear
* @Last Modified time: 2018-11-07 13:28:06
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

	//u := new(user)

	u := user{"bear", 19}

	//age := 18

	//fmt.Println(reflect.ValueOf(u))
	//
	//fmt.Println(reflect.ValueOf(&u))
	//
	//fmt.Println(reflect.ValueOf(&u).Elem())
	//
	//fmt.Println(reflect.ValueOf(age))

	defer func() {

		if p := recover(); p != nil {
			fmt.Println("paic is coming")

		}

	}()

	fmt.Println(reflect.ValueOf(u))

	ue := reflect.ValueOf(u).Elem()

	fmt.Println(ue.FieldByName("age"))

	//panic()

}
