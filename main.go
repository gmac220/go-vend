package main

import (
	"fmt"

	"github.com/gmac220/go-vend/vend"
)

func main() {
	var cewmachine = vend.Machine{
		Model:    "18239342",
		Address:  vend.NewAddress("101 Mitchell St", "Arlington", "TX", "76010"),
		Capacity: 100,
	}

	//fmt.Println(cewmachine)
	// fmt.Println("Capacity:", cewmachine.capacity)

	var p vend.Printer = cewmachine
	fmt.Println(p)
}
