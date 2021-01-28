package decorator

//素食
type VeggeMania struct {
}

//素食类价格
func (vm *VeggeMania) GetPrice() int {
	return 15
}
