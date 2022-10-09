package vm

import "fmt"

type ptEntry struct {
	sAddr int
	eAddr int
	name  string
}

var pt []ptEntry
var mem [0xFFFF]byte

func NewPt(psAddr int, peAddr int, pname string) {
	pt = append(pt, ptEntry{sAddr: psAddr, eAddr: peAddr, name: pname})
}

func MemCheck(addr int) { // checks if address violates a pointer from the pointer table
	for i, _ := range mem {
		fmt.Println(i)
	}
}

func Init() {
	for i, _ := range mem {
		mem[i] = 0
	}
}
