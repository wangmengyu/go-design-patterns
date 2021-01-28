package main

import "design_patterns/component"

func main() {

	//创建3个文件,分别赋值文件名
	file1 := &component.File{Name: "File1"}
	file2 := &component.File{Name: "File2"}
	file3 := &component.File{Name: "File3"}
	//创建1个文件夹, 赋予名字
	folder1 := &component.Folder{
		Name: "Folder1",
	}
	//将文件1加入到文件夹1
	folder1.Add(file1)
	//创建文件夹2, 赋予名字
	folder2 := &component.Folder{Name: "Folder2"}
	//将文件2, 3 加入到文件夹2
	folder2.Add(file2)
	folder2.Add(file3)
	//将文件夹1加入到文件夹2
	folder2.Add(folder1)
	//在文件夹2中搜索关键词rose
	folder2.Search("rose")

}
