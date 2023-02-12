package gun

type Ak47 struct {
	Gun // 包含了基本枪支所有属性
}

func NewAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
