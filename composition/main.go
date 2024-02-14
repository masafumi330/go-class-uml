package main

type Parent struct {
	name string
	Child
}

type Child struct {
	name string
}

func (c *Child) GetName() string {
	return c.name
}

func (p *Parent) GetName() string {
	return p.name
}

func main() {

	p := Parent{
		name: "Isao",
		Child: Child{
			name: "Taro",
		},
	}

	println(p.GetName())
	println(p.Child.GetName())
}
