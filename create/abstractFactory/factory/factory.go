package factory

import "fmt"

// ISportsFactory
// 抽象工厂模式：抽象工厂模式关注的是创建一组相关或相互依赖的对象。
// 它是一个工厂的集合，创建产品族中的多个对象。工厂里可以产出多种产品.
type ISportsFactory interface {
	MakeShoe() IShoe   //  鞋子
	MakeShirt() IShirt // 衣服
}

type Adidas struct {
}

func (a *Adidas) MakeShoe() IShoe {
	s := AdidasShoe{}
	s.SetLogo("adidas")
	s.SetSize(14)
	return &s

}
func (a *Adidas) MakeShirt() IShirt {
	s := AdidasShirt{}
	s.SetLogo("adidas")
	s.SetSize(14)
	return &s
}

type Nike struct {
}

func (a *Nike) MakeShoe() IShoe {
	s := NikeShoe{}
	s.SetLogo("nike")
	s.SetSize(14)
	return &s

}
func (a *Nike) MakeShirt() IShirt {
	s := NikeShirt{}
	s.SetLogo("nike")
	s.SetSize(14)
	return &s
}

type IShoe interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) GetLogo() string {
	return s.logo
}

func (s *Shoe) SetSize(size int) {
	s.size = size
}

func (s *Shoe) GetSize() int {
	return s.size
}

type AdidasShoe struct {
	Shoe
}
type NikeShoe struct {
	Shoe
}
type IShirt interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}
type Shirt struct {
	logo string
	size int
}

func (s *Shirt) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) SetSize(size int) {
	s.size = size
}

func (s *Shirt) GetSize() int {
	return s.size
}
func (s *Shirt) GetLogo() string {
	return s.logo
}

type AdidasShirt struct {
	Shirt
}

type NikeShirt struct {
	Shirt
}

func PrintShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
func PrintShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
