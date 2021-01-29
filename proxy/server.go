package proxy

//server主体
type Server interface {
	// 服务接口,  控制请求
	HandleRequest(url, method string) (int, string)
}
