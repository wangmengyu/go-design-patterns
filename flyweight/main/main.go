package main

import (
	"design_patterns/flyweight"
	"fmt"
)

func main() {
	g := flyweight.NewGame()
	_ = g.AddTerrorist()
	_ = g.AddTerrorist()
	_ = g.AddTerrorist()
	_ = g.AddTerrorist()
	_ = g.AddCounterTerrorist()
	_ = g.AddCounterTerrorist()
	_ = g.AddCounterTerrorist()

	dressFactoryIns := flyweight.GetDressFactorySingleInstance()
	for dressType, val := range dressFactoryIns.DressMap {
		fmt.Println(dressType, val.GetColor())
	}

	for _, v := range g.CounterTerrorist {
		fmt.Println(v)
	}

	for _, v := range g.Terrorist {
		fmt.Println(v)
	}

}
