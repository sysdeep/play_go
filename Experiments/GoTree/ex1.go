package main

import "log"

var PAYLOADEX1 []byte = []byte{
	0x11, 0x12, 0x13, 0x14,
	0x15, 0x16, 0x17, 0x18,
	0x11, 0x12, 0x13, 0x14,
	0x11, 0x12, 0x13, 0xaa,
}

func ex1() {
	log.Println("data count: ", len(PAYLOADEX1))

	tree := NewTree()

	for i := 0; i < (len(PAYLOADEX1) / 4); i++ {
		chunk := PAYLOADEX1[i*4 : i*4+4]
		tree.AppendChunk(chunk)
	}

	log.Println("tree count: ", tree.GetCount())

}
