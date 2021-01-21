package abstract_factory

type nike struct {
}

func (a *nike) MakeShoe() iShoe {
	return &nikeShore{
		shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (a *nike) MakeShirt() iShirt {
	return &nikeShirt{
		shirt{
			logo: "nike",
			size: 14,
		},
	}
}
