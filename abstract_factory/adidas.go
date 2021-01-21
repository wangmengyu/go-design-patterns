package abstract_factory

type adidas struct {
}

func (a *adidas) MakeShoe() iShoe {
	return &adidasShoe{
		shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *adidas) MakeShirt() iShirt {
	return &adidasShirt{
		shirt{
			logo: "adidas",
			size: 14,
		},
	}
}
