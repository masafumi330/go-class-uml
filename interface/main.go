package main

type Animal interface {
	GetName() string
	MakeNoise() string
}

type Dog struct {
	name string
}

func (d *Dog) GetName() string {
	return d.name
}

func (d *Dog) MakeNoise() string {
	return "Woof!"
}

type Cat struct {
	name string
}

func (c *Cat) GetName() string {
	return c.name
}

func (c *Cat) MakeNoise() string {
	return "Meow!"
}

func main() {
	animals := []Animal{
		&Dog{name: "Pochi"},
		&Cat{name: "Tama"},
	}

	for _, animal := range animals {
		println(animal.GetName(), "says", animal.MakeNoise())
	}
}
