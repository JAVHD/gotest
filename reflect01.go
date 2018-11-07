/*
* @Author: bear
* @Date:   2018-10-24 23:09:02
* @Last Modified by:   bear
* @Last Modified time: 2018-10-24 23:27:34
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var int x = 10
	//var a, b int = 10, 20
	var a int = 10
	var c *int = &a

	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(c))

	typec := reflect.TypeOf(c)

	fmt.Println(typec.Kind())

}
