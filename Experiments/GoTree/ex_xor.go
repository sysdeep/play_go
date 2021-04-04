package main

import "fmt"

/*
	((a ^ b) ^ c) ^ d = r

	d = ((a ^ b) ^ c) ^ r
	d^r = (a ^ b) ^ c



	нет данных для расшифровки...

*/
func exXor() {
	a := 12
	b := 14
	c := 32
	d := 15

	// x1 := a ^ b
	// x2 := x1 ^ c
	// x3 := x2 ^ d
	r := ((a ^ b) ^ c) ^ d

	fmt.Println("r: ", r)
	fmt.Println("d: ", d)
	fmt.Println("c: ", d^r)

	// fmt.Println("x1: ", x1)
	// fmt.Println("x2: ", x2)
	// fmt.Println("x3: ", x3)

	// y2 := x3 ^ d
	// fmt.Println("y2: ", y2)

	// fmt.Println("y1: ", y2^x3)
	// y1 := x2 ^ y2
	// a := x2 ^ y2

}
