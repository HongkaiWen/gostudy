package mymath

import "errors"

func Add(a int, b int) (c int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("a can't less than 0")
		return
	}
	return a + b, nil

}
