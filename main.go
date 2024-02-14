package main

import "fmt"

type Parent struct {
	Child
	Name string
}

type Child struct {
	Name string
}

func (c *Child) Hello() {
	fmt.Printf("My name is %s\n", c.Name)
}

func main() {
	p := Parent{
		Child: Child{Name: "Child"},
		Name:  "Parent",
	}
	p.Hello()
	p.Child.Hello()
}
