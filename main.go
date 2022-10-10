package main

import (
	"arrow/vm"
	"fmt"
)

func main() {
	fmt.Println("ArrowVM")
	vm.Init()
	fmt.Println(vm.MemCheck(11))
}
