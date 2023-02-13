package pizza

type IPizza interface { // pizza 抽象接口
	GetPrice() int // 获得总价
}

type VeggeMania struct {
}

func (p *VeggeMania) GetPrice() int {
	return 15 // 基本价格.
}

type TomatoTopping struct {
	Pizza IPizza // 顶盖会包含一个基本的pizza
}

func (c *TomatoTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice() // 价格是在原有基础上叠加
	return pizzaPrice + 7
}

type CheeseTopping struct {
	Pizza IPizza
}

func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 10
}
