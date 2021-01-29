package command

//命令接口. 可以被很多种命令来继承实现, 其中包含一个执行的方法需要被实现
type Command interface {
	Execute()
}
