package component

import "fmt"

//file 结构, 具有name属性
type File struct {
	Name string
}

//具有两个方法. 搜索关键字, 和获得文件名
func (f *File) Search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.Name)
}
func (f *File) GetName() string {
	return f.Name
}
