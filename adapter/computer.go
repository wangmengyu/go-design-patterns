package adapter

type IComputer interface { // 计算机接口. 可以被实现. 有一个插入到lightning接口的方法
	InsertIntoLightningPort()
}
