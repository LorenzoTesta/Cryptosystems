package main

import (
	"fmt"
	"../../../src/attack/classical/shift"
)

func main () {

	fmt.Println("********** ES 01 **********")

	c := "ycvejqwvhqtdtwvwu"

	fmt.Println("ciphertext:", c)

	for _, v := range shift.Exhaustive(c) {
		fmt.Println(v)
	}

	fmt.Println("********** ES 02 **********")

	shift.LetterFrequency(c)
}
