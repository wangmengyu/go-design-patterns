package main

import (
	"design_patterns/create/prototype/prototype"
	"fmt"
)

func main() {
	c1 := &prototype.Circle{
		X:      1,
		Y:      2,
		Radius: 3,
	}

	c2 := c1.Clone()
	fmt.Println(c2)

}
