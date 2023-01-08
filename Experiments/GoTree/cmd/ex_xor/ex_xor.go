package main

import "fmt"

/*
	((a ^ b) ^ c) ^ d = r

	d = ((a ^ b) ^ c) ^ r
	d^r = (a ^ b) ^ c



	нет данных для расшифровки...

	y1	y2

	a 	b		x1
	c 	d		x2

	x1 = a^b
	x2 = c^d
	y1 = a^c
	y2 = b^d

	x1^y1 = x2^y2 = b^c
	x1^y2 = x2^y1 = a^d

	a = x1^b = x1^y2^d = x1^y2^x2^c = x1^y2^x2^y1^a
	b = y2^d
	d = x2^c
	c = y1^a



*/
func exXor() {
	a := 12
	b := 14
	c := 32
	d := 15

	x1 := a ^ b
	x2 := c ^ d

	y1 := a ^ c
	y2 := b ^ d
	// x3 := x2 ^ d
	// r := ((a ^ b) ^ c) ^ d

	fmt.Println("x1: ", x1)
	fmt.Println("x2: ", x2)
	fmt.Println("y1: ", y1)
	fmt.Println("y2: ", y2)

	fmt.Println("x1y1: ", x1^y1)
	fmt.Println("x2y2: ", x2^y2)

	fmt.Println("b: ", x1^a)
	fmt.Println("c: ", y1^a)
	fmt.Println("d: ", x2^y1^a)
	// fmt.Println("y2: ", y2)

	// fmt.Println("x1: ", x1)
	// fmt.Println("x2: ", x2)
	// fmt.Println("x3: ", x3)

	// y2 := x3 ^ d
	// fmt.Println("y2: ", y2)

	// fmt.Println("y1: ", y2^x3)
	// y1 := x2 ^ y2
	// a := x2 ^ y2

	// data := [4][4]byte{
	// 	{12, 12, 23, 44},
	// 	{44, 55, 22, 10},
	// 	{55, 44, 33, 22},
	// 	{66, 77, 88, 22},
	// }

	// xs := make([]byte, 4)
	// ys := make([]byte, 4)

	// for _, x := range data{

	// 	for _, y :range x{
	// 		ys = append(ys, y)
	// 	}
	// }

}
