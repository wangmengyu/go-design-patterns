package builder

// Computer 分步骤构建结构体内容.
type Computer struct {
	CPU string
	RAM string
	HDD string
	GPU string
}

type IBuilder interface {
	SetCPU(string) IBuilder
	SetRAM(string) IBuilder
	SetHDD(string) IBuilder
	SetGPU(string) IBuilder
	Build() *Computer //  返回具体生成的整体结构体
}

type ComputerBuilder struct {
	computer *Computer // 具体实现构建器的结构
}

func (c *ComputerBuilder) SetCPU(cpu string) IBuilder {
	c.computer.CPU = cpu
	return c
}
func (c *ComputerBuilder) SetRAM(ram string) IBuilder {
	c.computer.RAM = ram
	return c
}
func (c *ComputerBuilder) SetHDD(hdd string) IBuilder {
	c.computer.HDD = hdd
	return c
}
func (c *ComputerBuilder) SetGPU(gpu string) IBuilder {
	c.computer.GPU = gpu
	return c
}
func (c *ComputerBuilder) Build() *Computer {
	return c.computer // 把本体内部的computer 返回出去
}
func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{
		computer: &Computer{}, // 初始化好内部的computer对象.
	}
}
