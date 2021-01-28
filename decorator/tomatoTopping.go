package decorator

// 番茄浇头
type TomatoTopping struct {
	//必须在某个Pizza上面
	Pizza Pizza
}

func (c *TomatoTopping) GetPrice() int {
	return c.Pizza.GetPrice() + 7 // 价格是Pizza价格+7
}
