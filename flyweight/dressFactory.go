package flyweight

import "fmt"

// 享元工厂

//公用常量
const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

type DressFactory struct { // 衣服工厂
	DressMap map[string]Dress // 内包含一个衣服的Map
}

var DressFacSingleInstance = &DressFactory{
	DressMap: make(map[string]Dress),
}

//获得衣服, 根据类型
func (d *DressFactory) GetDressByType(dressType string) (Dress, error) {
	if _, ok := d.DressMap[dressType]; ok == true {
		//已经创建了, 直接返回
		return d.DressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		//创建一个T类型的
		d.DressMap[dressType] = NewTerroristDress()
		return d.DressMap[dressType], nil
	}

	if dressType == CounterTerroristDressType {
		d.DressMap[dressType] = NewCounterTerroristDress()
		return d.DressMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

func GetDressFactorySingleInstance() *DressFactory {
	return DressFacSingleInstance
}
