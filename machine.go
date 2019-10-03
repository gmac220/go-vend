package main

/*
This is a struct representing a vending machines.
It stores a model, an address, and a capacity.
It has a String() method to override the normal
print behavior.
*/
type machine struct {
	model    string
	address  address
	capacity int
}

//This is a method to override normal print
func (m machine) String() string {
	return "This machine is located at\n" + m.address.String()
}
