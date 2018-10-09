/*
* @Author: bear
* @Date:   2018-09-11 10:35:34
* @Last Modified by:   bear
* @Last Modified time: 2018-09-11 10:39:55
 */

package main

import (
	"fmt"
)

func main() {

	//Sprintf 返回的是一个字符串
	//fmt.Sprintf(format, ...)

	i := []int{1, 2}

	s := fmt.Sprintf("v%", i)

	fmt.Println(s)

}
