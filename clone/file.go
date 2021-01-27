package inode

import "fmt"

type File struct {
	//文件名
	Name string //
}

func (f *File) Print(intention string) {
	fmt.Println(intention + f.Name)
}
func (f *File) Clone() INode {
	// 复制一个File对象出来,并且指定名称是clone出来的
	return &File{
		Name: f.Name + "_clone",
	}
}
