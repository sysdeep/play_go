package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")

	err, config := ConfigFromArgs()

	if err != nil {
		panic(err)
	}

	fmt.Println(config.FilePath)

	parse(config)
}
