package proxy

type Application struct {
}

func (a *Application) HandleRequest(url, method string) (int, string) {
	//只允许一下两种method的的请求
	if url == "/app/status" && method == "GET" {
		return 200, "ok"
	}

	if url == "/create/user" && method == "POST" {
		return 200, "ok"
	}

	return 404, "not found"

}
