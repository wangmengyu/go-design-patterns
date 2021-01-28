package component

import "fmt"

type Folder struct { // 组合: 组件集合
	components []Component // 组件列表
	Name       string
}

//搜索关键词
func (f *Folder) Search(keyword string) {
	fmt.Printf("Search recursively for keyword %s in folder %s\n", keyword, f.Name)
	for _, v := range f.components {
		v.Search(keyword)
	}
}

//添加组件到当前组件集合中
func (f *Folder) Add(c Component) {
	f.components = append(f.components, c)
}
