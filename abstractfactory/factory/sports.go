package factory

import "fmt"

type ISportsFactory interface {
	MakeShoe() IShoe
	// makeShort() IShirt
}

func NewSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	if brand == "nike" {
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("wrong brand type passed")
}
