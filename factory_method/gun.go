package factory_method

import "fmt"

type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}
func (g *gun) setPower(power int) {
	g.power = power
}
func (g *gun) getName() string {
	return g.name
}
func (g *gun) getPower() int {
	return g.power
}

func PrintDetails(g iGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Powner: %d\n", g.getPower())
}
