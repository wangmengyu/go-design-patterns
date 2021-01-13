package factory_method

type Ak47 struct {
	gun
}

func newAk47() iGun {
	return &Ak47{gun{
		name:  "Ak47 gun",
		power: 4,
	}}
}
