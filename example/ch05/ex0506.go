package main

import (
	"errors"
	"fmt"
)

func divAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		return numerator, denominator, errors.New("0で割ることはできません")
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

func ex0506() {
	x, y, z := divAndRemainder(5, 2)
	fmt.Println(x, y, z)
}
