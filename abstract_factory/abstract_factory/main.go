package main

import (
	"design_patterns/abstract_factory"
)

func main() {

	adidasFactory, _ := abstract_factory.GetSportsFactory("adidas")
	nikeFactory, _ := abstract_factory.GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	abstract_factory.PrintShoeDetail(nikeShoe)
	abstract_factory.PrintShirtDetails(nikeShirt)
	abstract_factory.PrintShoeDetail(adidasShoe)
	abstract_factory.PrintShirtDetails(adidasShirt)

}
