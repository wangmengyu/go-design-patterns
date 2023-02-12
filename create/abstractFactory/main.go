package main

import "design_patterns/create/abstractFactory/factory"

func main() {
	adidasFactory, _ := factory.GetSportsFactory("adidas")
	nikeFactory, _ := factory.GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	factory.PrintShoeDetails(nikeShoe)
	factory.PrintShirtDetails(nikeShirt)

	factory.PrintShoeDetails(adidasShoe)
	factory.PrintShirtDetails(adidasShirt)
}
