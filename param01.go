package main

import (
	"fmt"
)

func MyFunc(a ...int) {
	for i := range a {
		a[i] += 100
	}
}
