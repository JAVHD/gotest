/*
* @Author: bear
* @Date:   2018-10-27 22:49:23
* @Last Modified by:   bear
* @Last Modified time: 2018-10-27 22:59:44
 */
package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 100

	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(&a))
}
