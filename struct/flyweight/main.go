package main

import (
	"design_patterns/struct/flyweight/dress"
	"fmt"
)

/*
*
享元是一种结构型设计模式， 它允许你在消耗少量内存的情况下支持大量对象。
模式通过共享多个对象的部分状态来实现上述功能。
换句话来说， 享元会将不同对象的相同数据进行缓存以节省内存。
*/
func main() {
	game := dress.NewGame()

	//Add Terrorist
	game.AddTerrorist(dress.TerroristDressType)
	game.AddTerrorist(dress.TerroristDressType)
	game.AddTerrorist(dress.TerroristDressType)
	game.AddTerrorist(dress.TerroristDressType)

	//Add CounterTerrorist
	game.AddCounterTerrorist(dress.CounterTerrroristDressType)
	game.AddCounterTerrorist(dress.CounterTerrroristDressType)
	game.AddCounterTerrorist(dress.CounterTerrroristDressType)

	dressFactoryInstance := dress.GetDressFactorySingleInstance()

	for dressType, d := range dressFactoryInstance.DressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, d.GetColor())
	}
}
