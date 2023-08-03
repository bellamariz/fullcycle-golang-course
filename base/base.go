package base

import "fmt"

// foundation for more complex data types, such as structs

type Client struct {
	Name    string
	Age     int
	Active  bool
	Addr    Address // new struct field of type Address
	Address         // struct composition
}

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Company struct {
	Name   string
	Active bool
}

// structs have methods

func (cp *Company) Deactivate() {
	cp.Active = false
}

func (c *Client) Deactivate() {
	c.Active = false
}

// interface - if structs have the same method, they can be defined through the interface that implements said method

type Person interface {
	Deactivate()
}

func Deactivate(p Person) {
	p.Deactivate()
}

// test program

func Structs() {
	patrick := Client{
		Name:   "Patrick Stump",
		Age:    32,
		Active: true,
	}

	falloutboy := Company{
		Name:   "Fall Out Boy",
		Active: true,
	}

	fmt.Println(patrick, falloutboy)

	patrick.Address.Street = "abbey road"
	patrick.Addr.Street = "cornelia street"

	Deactivate(&patrick)
	Deactivate(&falloutboy)

	fmt.Println(patrick, falloutboy)
}

func sum(a, b int) int {
	a = 50
	return a + b
}

func sumP(a, b *int) int {
	*a = 50
	return *a + *b
}

// pointer - variable that contains the address to another variable in memory
func Pointers() {
	a := 10
	var p *int = &a

	*p = 20

	fmt.Printf("variable: address %p - value %d\n", &a, a)
	fmt.Printf("pointer: address %p - value %p - reference %d\n", &p, p, *p)

	// when passed as arguments to functions, variables are only copied
	// but when a pointer is passed as argument, the variable's real value is altered

	varA := 10
	varB := 10

	s := sum(varA, varB)
	fmt.Printf("var: %d - sum: %d\n", varA, s)

	sp := sumP(&varA, &varB)
	fmt.Printf("var: %d - sum: %d\n", varA, sp)
}
