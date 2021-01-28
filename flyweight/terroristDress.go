package flyweight

type TerroristDress struct {
	Color string
}

func (td *TerroristDress) GetColor() string {
	return td.Color
}

//创建一个TDress
func NewTerroristDress() *TerroristDress {
	return &TerroristDress{Color: "red"}
}
