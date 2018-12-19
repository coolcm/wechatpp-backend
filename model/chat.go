package model

import (
	"fmt"
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
func CreateChat(QuserId string, AuserId string) (chat Chat) {
	startTime := time.Now()
	s := QuserId + AuserId + startTime.String()
	hash := utils.CalHash(s)

	chat = Chat{
		Hash: hash,
		QuserId: QuserId,
		AuserId: AuserId,
		StartTime: startTime,
		Grade: -1,  // 用于是否完成打分的逻辑判断
	}

	Db.Create(&chat)
	return
}

// 查询答疑记录
func QueryChat(hash string) (Chat) {
	var chat Chat
	Db.Where("hash = ?", hash).First(&chat)
	return chat
}

// 记录答疑结束
func (chat *Chat) End() {
	if chat.EndTime.Before(chat.StartTime) {
		chat.EndTime = time.Now()
		Db.Save(chat)
	}
}

// 给本次答疑打分
func (chat *Chat) Score(grade int){

	// 若答疑未结束，则将其设为结束
	if chat.EndTime.Before(chat.StartTime) {
		chat.EndTime = time.Now()
	}
	// 若该答疑已评过分，则直接返回
	if chat.Grade >= 0 {
		return
	}
	// 自动计算生成答疑所需支付的token
	chatTime := chat.EndTime.Sub(chat.StartTime)
	fmt.Println(chatTime.Minutes())
	token := rand.Intn(10 + grade/10 + int(chatTime.Minutes()))
	fmt.Println(token)
	chat.Grade = grade
	chat.Token = token
	Db.Save(chat)

	// 更新提问者和答疑者的相关信息
	UpdateUserQATime(Db, chat.QuserId, chat.AuserId, chat.EndTime.Sub(chat.StartTime))
	SetUserPendingToken(Db, chat.QuserId, chat.Token)
	return
}