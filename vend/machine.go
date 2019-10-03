package vend

/*
Machine is a struct representing a vending machines.
It stores a model, an address, and a capacity.
It has a String() method to override the normal
print behavior.
*/
type Machine struct {
	Model    string
	Address  *Address
	Capacity int
}

//This is a method to override normal print
func (m Machine) String() string {
	return "This machine is located at\n" + m.Address.String()
}
