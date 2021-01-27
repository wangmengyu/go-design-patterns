package inode

import "fmt"

type Folder struct {
	Children []INode // 子文件, 一系列的
	Name     string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name) // 打印一下名字
	//并且打印所有子文件夹的名字
	for _, v := range f.Children {
		v.Print(indentation + indentation) //子目录前要多一个缩进
	}
}

func (f *Folder) Clone() INode {
	//复制目录本身 以及它下面的子目录
	cloneFolder := &Folder{
		Name: f.Name + "_clone",
	}
	tmpCh := make([]INode, 0)
	for _, v := range f.Children {
		cpy := v.Clone() // 将子目录复制出来. 复制给临时变量
		tmpCh = append(tmpCh, cpy)
	}
	cloneFolder.Children = tmpCh
	return cloneFolder
}
