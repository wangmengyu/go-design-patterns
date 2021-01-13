package factory_method

type Ak47 struct {
	gun
}

func newAk47() iGun { // 工厂方法, 返回接口实例化的类
	return &Ak47{gun{
		name:  "Ak47 gun",
		power: 4,
	}}
}
