package component

import "fmt"

type File struct {
	Name string
}

func (f *File) Search(keyword string) {
	// 搜索自己
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.Name)
}

func (f *File) GetName() string {
	return f.Name
}
