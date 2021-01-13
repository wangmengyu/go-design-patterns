package factory_method

//定义接口
type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
