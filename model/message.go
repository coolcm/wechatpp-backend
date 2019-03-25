package model

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

// 聊天消息数据表
type Message struct {
	Id           uint
	FromWechatId string    //消息发送者微信号
	ToWechatId   string    //消息接收者微信号
	CreateTime   time.Time //消息创建时间
	Content      string    //消息内容
	Status       bool      //是否已读
}

// 创建新消息
func CreateMessage(FromWechatId string, ToWechatId string, Content string) {
	var Message = Message{
		FromWechatId: FromWechatId,
		ToWechatId: ToWechatId,
		Content: Content,
		Status: false,
		CreateTime: time.Now(),
	}
	Db.Create(&Message)
}

// 获取用户所有未读数据
func GetUnreadMessage(ToWechatId string) (messages []Message){
	Db.Where("to_wechat_id = ? AND Status = false", ToWechatId).Find(&messages)
	return
}

// 消息已读后更新数据库消息状态
func UpdateMessageStatus(message Message) {
	Db.First(&message).Update("status", true)
}
