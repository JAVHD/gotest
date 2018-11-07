/*
* @Author: bear
* @Date:   2018-10-27 22:03:53
* @Last Modified by:   bear
* @Last Modified time: 2018-10-27 22:12:42
 */
package main

import (
	"fmt"
	"reflect"
)

func main() {

	type A int

	var a int = 1

	var b A = 123

	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))

	fmt.Println(reflect.TypeOf(b).Kind())

	//reflect.Ptr

}
