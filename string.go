/*
* @Author: bear
* @Date:   2018-12-04 10:55:56
* @Last Modified by:   bear
* @Last Modified time: 2018-12-04 11:14:23
 */

package main

import (
	"fmt"
	"strings"
)

func main() {

	port := ":8090"

	test := strings.Split(port, ":")

	test01 := []string{"hello", "world"}

	fmt.Println(len(test))
	fmt.Println(test)
	fmt.Println(test01)
	fmt.Printf("%v", test)

}
