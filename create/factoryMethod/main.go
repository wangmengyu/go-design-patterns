package main

import "design_patterns/create/factoryMethod/gun"

func main() {
	ak47, _ := gun.GetGun("ak47")
	musket, _ := gun.GetGun("musket")

	gun.PrintDetails(ak47)
	gun.PrintDetails(musket)
}
