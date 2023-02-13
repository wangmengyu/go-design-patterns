package dress

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func NewPlayer(playerType, dressType string) *Player {
	dress, _ := GetDressFactorySingleInstance().GetDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress, // 其实是从map里获得的成员值
	}
}

func (p *Player) NewLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
