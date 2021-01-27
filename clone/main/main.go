package main

import (
	inode "design_patterns/clone"
	"fmt"
)

func main() {
	//定义三个文件
	file1 := &inode.File{Name: "File1"}
	file2 := &inode.File{Name: "File2"}
	file3 := &inode.File{Name: "File3"}
	//定义一个目录,目录里本来包含第一个文件
	folder1 := &inode.Folder{
		Children: []inode.INode{file1},
		Name:     "Folder1",
	}

	//定义第二个目录, 里面包含目录1和文件2,文件3
	folder2 := &inode.Folder{
		Children: []inode.INode{folder1, file2, file3},
		Name:     "Folder2",
	}

	//打印第二个目录本身
	fmt.Println("Printing for Folder2")
	folder2.Print("  ")

	//打印第二个目录的clone版本
	cloneFolder := folder2.Clone()
	fmt.Println("Printing for Clone Folder2")
	cloneFolder.Print("  ")

}
