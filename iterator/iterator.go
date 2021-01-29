package iterator

type Iterator interface { // 迭代器接口, 有两个方法, 检查是否有下一个, 还有获得下一个
	HasNext() bool
	GetNext() interface{}
}
