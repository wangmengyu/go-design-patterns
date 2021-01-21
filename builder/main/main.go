package main

import (
	"design_patterns/builder"
	"fmt"
)

func main() {
	normalBuilder := builder.GetBuilder("normal")
	iglooBuilder := builder.GetBuilder("igloo")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Println("Normal House Door Type:", normalHouse.DoorType)
	fmt.Println("Normal House Window Type :", normalHouse.WindowType)
	fmt.Println("Normal House Num Floor :", normalHouse.Floor)

	director = builder.NewDirector(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Println("igloo House Door Type:", iglooHouse.DoorType)
	fmt.Println("igloo House Window Type :", iglooHouse.WindowType)
	fmt.Println("igloo House Num Floor :", iglooHouse.Floor)

}
