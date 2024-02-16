package main

import (
	"fmt"
	"go-class-uml/abstractfactory/factory"
)

func main() {
	nike, err := factory.NewSportsFactory("nike")
	if err != nil {
		panic(err)
	}

	shoe := nike.MakeShoe()
	shoe.SetLogo("nike")
	shoe.SetSize(12)
	fmt.Printf("Logo: %s\n", shoe.GetLogo())
	fmt.Printf("Size: %d\n", shoe.GetSize())

}
