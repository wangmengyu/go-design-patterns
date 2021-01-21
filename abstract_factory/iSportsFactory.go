package abstract_factory

import "fmt"

//一个运动装备接口,
type iSportsFactory interface {
	MakeShoe() iShoe   //做鞋子
	MakeShirt() iShirt // 做衬衫
}

/**
获得运动工厂 接口, 按照品牌
*/
func GetSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	} else if brand == "nike" {
		return &nike{}, nil
	}

	return nil, fmt.Errorf("WRONG brand type")
}
