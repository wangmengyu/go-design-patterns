package main

import "design_patterns/factory_method"

func main() {
	ak47, _ := factory_method.GetGun("ak47")
	musket, _ := factory_method.GetGun("musket")
	factory_method.PrintDetails(ak47)
	factory_method.PrintDetails(musket)
}
