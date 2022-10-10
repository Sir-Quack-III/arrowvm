package vm

import (
	"fmt"
	"os"
)

type Type int

const (
	i64 Type = iota
	i32
	i16
	i8
	f64
	f32
	char
	str
	bl
)

type ptEntry struct {
	sAddr int
	eAddr int
	name  string
	tp    Type
}

var pt []ptEntry
var mem [0xFFFF]byte

func Panic(title string, desc string, end bool) {
	fmt.Print(title + ": ")
	fmt.Println(desc)
	if end {
		os.Exit(1)
	}
}

func NewPt(psAddr int, peAddr int, pname string, ptp Type) {
	pt = append(pt, ptEntry{sAddr: psAddr, eAddr: peAddr, name: pname, tp: ptp})
}

func Alloc(tp Type, len int, name string) {
	for i, _ := range mem {
		if MemCheck(i) && MemCheck(i+len) {
			NewPt(i, i+len, name, tp)
			return
		}
	}

	Panic("MEMLEAK", "No memory left to allocate!", true)
}

func MemCheck(addr int) bool { // checks if address violates a pointer from the pointer table
	for _, p := range pt {
		for i := p.sAddr; i <= p.eAddr; i++ {
			if i == addr {
				return true
			}
		}
	}

	return false
}

func Init() {
	for i, _ := range mem {
		mem[i] = 0
	}
}
