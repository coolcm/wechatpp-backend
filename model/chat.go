package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sjtucsn/wechatpp-backend/utils"
	"math/rand"
	"time"
)

// 用户答疑数据表
type Chat struct {
	Id        uint
	Hash      string    // 通过hash256计算得出的唯一哈希值作为答疑的唯一认证
	QuserId   string    // 提问用户的id
	AuserId   string    // 答疑用户的id
	StartTime time.Time // 本次答疑开始时间
	EndTime   time.Time // 本次答疑结束时间
	Grade     int       // 提问者对本次答疑的评价（分数0~10）
	Token     int       // 本次答疑需要支付的token数
	State     bool      // 是否本次答疑已完成（包括token已支付）
}

// 增加一条答疑记录
func CreateChat(db *gorm.DB, QuserId string, AuserId string) (chat Chat) {
	startTime := time.Now()
	s := QuserId + AuserId + startTime.String()
	hash := utils.CalHash(s)

	chat = Chat{
		Hash: hash,
		QuserId: QuserId,
		AuserId: AuserId,
		StartTime: startTime,
	}

	db.Create(&chat)
	return
}

// 记录答疑已完成
func EndChat(db *gorm.DB, hash string) (Chat) {
	var chat Chat
	db.Where("hash = ?", hash).First(&chat).Update("end_time", time.Now())
	return chat
}

// 给本次答疑打分
func ScoreChat(db *gorm.DB, hash string, grade int) (Chat) {
	var chat Chat
	db.Where("hash = ?", hash).First(&chat)
	if chat.Id == 0 {
		return chat
	}
	chatTime := chat.EndTime.Sub(chat.StartTime)
	fmt.Println(chatTime.Minutes())
	token := rand.Intn(10 + grade/10 + int(chatTime.Minutes()))
	fmt.Println(token)
	chat.Grade = grade
	chat.Token = token
	db.Save(&chat)
	return chat
}