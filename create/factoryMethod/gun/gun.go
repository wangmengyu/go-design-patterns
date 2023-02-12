package gun

import "fmt"

// IGun 工厂方法模式关注的是创建一个单独的对象。它为每种产品定义一个工厂，创建一种单独的产品。
type IGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}

type Gun struct { // 枪支具有统一的属性.
	name  string
	power int
}

func (g *Gun) SetName(name string) {
	g.name = name
}

func (g *Gun) GetName() string {
	return g.name
}

func (g *Gun) SetPower(power int) {
	g.power = power
}

func (g *Gun) GetPower() int {
	return g.power
}

func GetGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return NewAk47(), nil
	}
	if gunType == "musket" {
		return NewMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

func PrintDetails(g IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
