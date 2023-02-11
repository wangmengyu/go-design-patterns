package collection

import "design_patterns/behavioral/iterator/iterator"

type Collection interface {
	CreateIterator() iterator.Iterator // 创建迭代器
}
