package main

import (
	"gotree/internal/LZTree"
	"log"
)

func main() {
	log.Println("Start app")

	//--- LZTree --------------------------------------------------------------
	// LZTree.ExampleLZTree1() // fake data
	// LZTree.ExampleLZTree2() // file data
	// LZTree.ExampleLZTree3() // calgary file data
	// LZTree.ExampleLZTree4() // all calgary file data
	LZTree.ExampleLZTree5() // all calgary file data in 1 tree
	//--- LZTree --------------------------------------------------------------

	log.Println("End app")
}
