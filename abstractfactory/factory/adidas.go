package factory

type Adidas struct{}

func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShoe{}
}

type AdidasShoe struct {
	Shoe
}
