package main

import (
	"design_patterns/struct/decorator/pizza"
	"fmt"
)

func main() {

	p := &pizza.VeggeMania{}

	//Add cheese topping
	pizzaWithCheese := &pizza.CheeseTopping{
		Pizza: p,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &pizza.TomatoTopping{
		Pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n",
		pizzaWithCheeseAndTomato.GetPrice())

}
