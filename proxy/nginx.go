package proxy

//nginx 结构, 包含 应用, 最大请求数, 频率限制map
type Nginx struct {
	Application       *Application
	MaxAllowedRequest int
	RateLimiter       map[string]int
}

//创建新的nginx服务
func NewNginxServer() *Nginx {
	return &Nginx{
		Application:       &Application{},
		MaxAllowedRequest: 2,
		RateLimiter:       make(map[string]int),
	}
}

//实现Server的接口方法,

func (n *Nginx) HandleRequest(url, method string) (int, string) {
	//看当前URL访问数量是不是大于设置的值了
	if n.CheckRateLimit(url) == false {
		return 403, "Over max limit"
	}

	//没有超过数量, 再去使用应用的请求处理
	return n.Application.HandleRequest(url, method)

}

/**
检查当前url 的访问rate,
*/
func (n *Nginx) CheckRateLimit(url string) bool {
	//先判断当前url有没有请求数
	if _, ok := n.RateLimiter[url]; ok == false {
		//没有, 初始化设置成1
		n.RateLimiter[url] = 1
	}

	//检查当前url的请求数量是不是大于最大访问值
	if n.RateLimiter[url] > n.MaxAllowedRequest {
		return false
	}

	//没有大于. 做叠加操作
	n.RateLimiter[url] += 1
	return true

}
