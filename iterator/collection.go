package iterator

type Collection interface {
	//集合的基类接口, 有一个获得迭代器的方法
	CreateIterator() Iterator // 返回迭代器出来
}
