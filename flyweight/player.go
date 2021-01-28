package flyweight

type Player struct {
	Dress      Dress
	PlayerType string
	Lat        int
	Long       int
}

// 构建新的角色
func NewPlayer(playerType, dressType string) (*Player, error) {
	dress, _ := GetDressFactorySingleInstance().GetDressByType(dressType)
	return &Player{
		Dress:      dress,
		PlayerType: playerType,
	}, nil
}

func (p *Player) NewLocation(lat, long int) {
	p.Lat = lat
	p.Long = long
}
