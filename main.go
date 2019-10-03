package main

import "fmt"

func main() {
	var cewmachine = machine{
		model: "18239342",
		address: address{
			street: "101 Mitchell St",
			city:   "Arlington",
			state:  "TX",
			zip:    "76010",
		},
		capacity: 100,
	}

	//fmt.Println(cewmachine)
	// fmt.Println("Capacity:", cewmachine.capacity)

	var p printer = cewmachine
	var a printer = cewmachine.address
	fmt.Println(p)
	fmt.Println(a)
}
