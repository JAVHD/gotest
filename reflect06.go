/*
* @Author: bear
* @Date:   2018-11-06 14:52:20
* @Last Modified by:   bear
* @Last Modified time: 2018-11-06 23:30:03
 */
package main

import (
	"fmt"
	"reflect"
)

func main() {

	type user struct {
		name string `field:"name" type:"varchar(50)"`
		age  int    `field:"age" type:"int"`
	}

	//u := user{"bear", 18}
	u := &user{"bear", 18}

	//fmt.Println(reflect.TypeOf(user{"bear", 18}))
	fmt.Println(reflect.TypeOf(u))

	ut := reflect.TypeOf(u).Elem()

	//fmt.Println(ut)

	for i := 0; i < ut.NumField(); i++ {
		fmt.Println(ut.Field(i).Tag)
	}

}
