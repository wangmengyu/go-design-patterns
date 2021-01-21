package abstract_factory

import "fmt"

// 鞋子的接口.
type iShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

//鞋子的实例
type shirt struct {
	logo string
	size int
}

func (s *shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *shirt) setSize(size int) {
	s.size = size
}

func (s *shirt) getLogo() string {
	return s.logo
}

func (s *shirt) getSize() int {
	return s.size
}

func PrintShirtDetails(s iShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
