# 华为消息推送

Golan版华为消息推送SDK

# 示例

```
package main

import (
	"fmt"
	hw "github.com/dongjialong2006/hw-push"
)

func main() {
	client := hw.NewClient("1", "***")	// 注册的客户端ID和密钥

	tokens := []string{"123", "234"} 	// 设备对应的Token
	msg := hw.NewMessage()
	msg.SetContent("Hello World") 		// 推送的消息
	msg.SetTitle(req.Subject)     		// 推送消息的主题
		
	resp, err := client.Push(tokens,msg.Json())
	if nil != err {
		fmt.Println(err)
		return
	}
	
	fmt.Println("resp", resp)
}

```