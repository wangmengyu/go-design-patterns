package decorator

type CheeseTopping struct {
	//也必须包含Pizza
	Pizza Pizza
}

func (ct *CheeseTopping) GetPrice() int {
	return ct.Pizza.GetPrice() + 10 // pizza价格加上10
}
