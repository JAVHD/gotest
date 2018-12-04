/*
* @Author: bear
* @Date:   2018-12-04 14:22:41
* @Last Modified by:   bear
* @Last Modified time: 2018-12-04 14:27:54
 */

package main

import (
	"fmt"
)

//append用来将元素添加到切片末尾并返回结果
func main() {

	s1 := []int{1, 2, 3, 4}

	s2 := append(s1, 7, 8, 9)

	fmt.Printf("%v", s1)
	fmt.Printf("%v", s2)

}
