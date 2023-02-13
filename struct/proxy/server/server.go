package server

type Server interface {
	HandleRequest(string, string) (int, string) // 控制请求
}

type Nginx struct { // 具体实现server的结构
	Application       *Application   // 代理要包含需要请求的真实应用成员
	MaxAllowedRequest int            // 最大访问请求
	RateLimiter       map[string]int // 访问频率限制
}

func NewNginxServer() *Nginx {
	return &Nginx{
		Application:       &Application{},
		MaxAllowedRequest: 2,
		RateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) HandleRequest(url, method string) (int, string) {
	/**
	代理模式是一种非常常用的设计模式，在软件开发中经常使用。
	它的主要优点在于可以在不改变目标对象的情况下控制对这个对象的访问，
	从而达到解耦、保护目标对象、提高系统灵活性等目的。
	代理模式在许多领域都有应用，例如在网络编程中使用代理服务器访问互联网，
	在虚拟技术中使用虚拟代理加载大型图像，在安全系统中使用保护代理控制对敏感对象的访问等。
	总的来说，代理模式是一种非常有用的设计模式，在各种不同的场景中都有应用，
	因此它是一种非常常用的设计模式。
	*/
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	//  没有问题在给到application 进一步处理请求
	return n.Application.HandleRequest(url, method)
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.RateLimiter[url] == 0 {
		n.RateLimiter[url] = 1
	}
	if n.RateLimiter[url] > n.MaxAllowedRequest { //  控制最大请求次数
		return false
	}
	n.RateLimiter[url] = n.RateLimiter[url] + 1
	return true
}
