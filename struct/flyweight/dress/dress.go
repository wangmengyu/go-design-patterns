package dress

import "fmt"

const (
	//TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	//CounterTerrroristDressType terrorist dress type
	CounterTerrroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		DressMap: make(map[string]Dress), // 全局的衣服map
	}
)

type DressFactory struct {
	DressMap map[string]Dress
}

func (d *DressFactory) GetDressByType(dressType string) (Dress, error) {
	if d.DressMap[dressType] != nil {
		// 如果已经定义过. 直接返回
		return d.DressMap[dressType], nil
	}

	// 按照不同类型创建不同类型的衣服,
	if dressType == TerroristDressType {
		d.DressMap[dressType] = NewTerroristDress()
		return d.DressMap[dressType], nil
	}
	if dressType == CounterTerrroristDressType {
		d.DressMap[dressType] = NewCounterTerroristDress()
		return d.DressMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

func GetDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}

type Dress interface {
	GetColor() string // 获得颜色
}
