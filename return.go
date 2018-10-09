package main

import (
	"errors"
	"fmt"
)

func div(x, y int) (int, error) {

	if y == 0 {
		return 0, errors.New("y is 0")
	}

	return x / y, nil

}

func main() {

	_, b := div(1, 0)

	if b != nil {
		fmt.Println(b)
	}
}
