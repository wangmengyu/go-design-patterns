package decorator

//Pizza 接口. 有一种方法叫获得价格
type Pizza interface {
	GetPrice() int // 返回的是整数的价格
}
