package dress

type TerroristDress struct {
	color string
}

func (t *TerroristDress) GetColor() string {
	return t.color
}

func NewTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"} // 红色
}
