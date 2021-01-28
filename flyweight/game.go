package flyweight

type Game struct {
	Terrorist        []Player
	CounterTerrorist []Player
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) AddTerrorist() error {
	p, err := NewPlayer("T", TerroristDressType)
	if err != nil {
		return err
	}
	g.Terrorist = append(g.Terrorist, *p)
	return nil
}
func (g *Game) AddCounterTerrorist() error {
	p, err := NewPlayer("CT", CounterTerroristDressType)
	if err != nil {
		return err
	}
	g.Terrorist = append(g.Terrorist, *p)
	return nil
}
