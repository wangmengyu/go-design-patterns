package component

type Component interface {
	//组件只有个搜索公用接口方法
	Search(keyword string)
}
