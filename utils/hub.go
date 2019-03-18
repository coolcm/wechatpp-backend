package utils

import "github.com/gorilla/websocket"

type Client struct {
	Hubs *Hub

	// The websocket connection.
	Conn *websocket.Conn

	// 在线聊天本方微信号
	WechatId string

	// 在线聊天对方微信号
	TargetWechatId string

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

// Hub maintains the set of clients
type Hub struct {
	Clients map[string]*Client
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[string]*Client),
	}
}

//func (h *Hub) run() {
//	for {
//		select {
//		case client := <-h.register:
//			h.clients[client] = true
//		case client := <-h.unregister:
//			if _, ok := h.clients[client]; ok {
//				delete(h.clients, client)
//				close(client.send)
//			}
//		case message := <-h.broadcast:
//			for client := range h.clients {
//				select {
//				case client.send <- message:
//				default:
//					close(client.send)
//					delete(h.clients, client)
//				}
//			}
//		}
//	}
//}
