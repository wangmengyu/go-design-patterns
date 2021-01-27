package adapter

import "fmt"

type Client struct {
	//客户端
}

//客户端能插入接口到电脑
func (c *Client) InsertLightningConnectorIntoComputer(com IComputer) {
	fmt.Println("Client inserts Lightning connector into computer")
	com.InsertIntoLightningPort() // 电脑, 插入到Lightning 接口里
}
