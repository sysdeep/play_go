package models

import "fmt"

type MainSupply struct {
	IsState bool
	IsError bool
	IsBlock bool
}

func (m *MainSupply) PowerOn() {
	fmt.Println("power on")
}
