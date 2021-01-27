package inode

type INode interface {
	Print(string) // 打印内容
	Clone() INode // 克隆接口
}
