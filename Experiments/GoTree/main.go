package main

import (
	"gotree/LZTree"
	"log"
)

func main() {
	log.Println("Start app")
	//--- ex1 -----------------------------------------------------------------
	// ex1()

	//--- ex_file -------------------------------------------------------------
	// exFile("./main.go", 4)
	// exFile("./main.go", 8)
	// exFile("./ex_file.go", 4)
	// exFile("/home/nika/screen_2021-03-25_19-16-53.mp4", 4)
	// exFile("/home/nika/screen_2021-03-25_19-16-53.mp4", 8)

	//--- ex xor --------------------------------------------------------------
	// exXor()

	//--- toffoli -------------------------------------------------------------
	// Toffoli()

	//--- LZTree --------------------------------------------------------------
	// LZTree.ExampleLZTree1() // fake data
	// LZTree.ExampleLZTree2() // file data
	// LZTree.ExampleLZTree3() // calgary file data
	LZTree.ExampleLZTree4() // all calgary file data
	//--- LZTree --------------------------------------------------------------

	log.Println("End app")
}
