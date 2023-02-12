package main

import (
	"design_patterns/create/builder/builder"
	"fmt"
)

func main() {
	com := builder.NewComputerBuilder().
		SetCPU("Intel Core i7").
		SetRAM("16GB").
		SetHDD("512GB SSD").
		SetGPU("Nvidia GeForce RTX 3080").
		Build()
	fmt.Println(com)
}
