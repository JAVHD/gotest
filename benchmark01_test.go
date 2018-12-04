/*
* @Author: bear
* @Date:   2018-11-09 09:28:16
* @Last Modified by:   bear
* @Last Modified time: 2018-11-09 09:28:38
 */

package benchmark01__test

import (
	//"fmt"
	"testing"
)


func add(x, y int) int {
	return x + y
}


func BenchmarkAdd(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = add(1, 2)
	}

}


