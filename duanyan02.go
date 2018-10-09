/*
* @Author: bear
* @Date:   2018-09-04 22:50:59
* @Last Modified by:   bear
* @Last Modified time: 2018-09-04 23:00:51
 */
package main

import (
	"fmt"
)

var Tester interface{}

func Test(a interface{}) string {

	if v, ok := a.(string); ok {
		fmt.Println(v, ok)
		return v
	}

	//return string(a)
	//return a
	return ""
}

func main() {
	//b := Test(10)
	b := Test("bear")

	fmt.Println(b)
}
