package webSocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sjtucsn/wechatpp-backend/utils"
	"net/http"
	"time"
)

type Client = utils.Client
type Hub = utils.Hub

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 处理ws请求
func WsHandler(c *gin.Context, hub *Hub) {
	weChatId := c.Query("wechat_id")
	targetWeChatId := c.Query("target_wechat_id")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &Client{
		Hubs: hub,
		Conn: conn,
		WechatId: weChatId,
		TargetWechatId: targetWeChatId,
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
	// 将该用户加入聊天Clients
	hub.Clients[weChatId + targetWeChatId] = client

	// 必须死循环，gin通过协程调用该handler函数，一旦退出函数，ws会被主动销毁
	for {
		_, reply, err := client.Conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(string(reply))
		// 查找聊天对像的Client
		targetClient := client.Hubs.Clients[targetWeChatId + weChatId]
		// 将发送的信息发给聊天对像
		if targetClient != nil {
			if err := targetClient.Conn.WriteMessage(1, reply); err != nil {
				fmt.Println(err)
			}
		} else {
			println("对方不在线")
		}
	}
}
