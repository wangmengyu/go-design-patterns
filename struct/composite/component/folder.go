package component

import "fmt"

type Folder struct {
	Components []Component // 文件夹下面可以有多个文件或者文件夹.
	Name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.Name)
	for _, composite := range f.Components {
		// 在当前内容列表内进行搜索
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c Component) {
	f.Components = append(f.Components, c) //追加到当前内容列表
}
