package main

import "design_patterns/struct/composite/component"

func main() {
	/**
	组合是一种结构型设计模式， 你可以使用它将对象组合成树状结构，
	并且能像使用独立对象一样使用它们。
	对于绝大多数需要生成树状结构的问题来说， 组合都是非常受欢迎的解决方案。
	组合最主要的功能是在整个树状结构上递归调用方法并对结果进行汇总。
	*/
	file1 := &component.File{Name: "File1"}
	file2 := &component.File{Name: "File2"}
	file3 := &component.File{Name: "File3"}

	folder1 := &component.Folder{
		Name: "Folder1",
	}

	folder1.Add(file1)

	folder2 := &component.Folder{
		Name: "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}
