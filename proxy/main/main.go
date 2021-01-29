package main

import (
	"design_patterns/proxy"
	"fmt"
)

func main() {
	//创建nginx
	nginxServer := proxy.NewNginxServer()
	//创建url1=/app/status
	appStatusUrl := "/app/status"
	//创建url2=/create/user
	createUserUrl := "/create/user"

	//请求
	code, body := nginxServer.HandleRequest(appStatusUrl, "GET")
	fmt.Printf("code= %d body=%s \n", code, body)

	code, body = nginxServer.HandleRequest(appStatusUrl, "GET")
	fmt.Printf("code= %d body=%s \n", code, body)

	//在请求会超过流量
	code, body = nginxServer.HandleRequest(appStatusUrl, "GET")
	fmt.Printf("code= %d body=%s \n", code, body)

	//请求另一个
	code, body = nginxServer.HandleRequest(createUserUrl, "POST")
	fmt.Printf("code= %d body=%s \n", code, body)

	//用错误的method
	code, body = nginxServer.HandleRequest(createUserUrl, "GET")
	fmt.Printf("code= %d body=%s \n", code, body)
}
