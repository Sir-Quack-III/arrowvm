package main

import (
	"arrow/vm"
	"fmt"
)

func main() {
	fmt.Println("ArrowVM")
	vm.Init()
	vm.NewPt(0, 10, "pointer")
}
