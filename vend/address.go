package vend

type Address struct {
	street string
	city   string
	state  string
	zip    string
}

//NewAddress is a Constructor for Address
func NewAddress(street string, city string, state string, zip string) *Address {
	n := Address{
		street: street,
		city:   city,
		state:  state,
		zip:    zip,
	}
	return &n
}

func (a Address) String() string {
	var output string
	output = a.street + "\n" + a.city + ", " + a.state + " " + a.zip + "\n"
	return output
}
