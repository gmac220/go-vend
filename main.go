package main

import "fmt"

type machine struct {
	model    string
	address  string
	capacity int
}

func main() {
	var cewmachine = machine{
		model:    "18239342",
		address:  "100 Mitchell St",
		capacity: 100,
	}

	fmt.Println(cewmachine)
}
