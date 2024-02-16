package factory

type Nike struct{}

func (a *Nike) MakeShoe() IShoe {
	return &NikeShoe{}
}

type NikeShoe struct {
	Shoe
}
