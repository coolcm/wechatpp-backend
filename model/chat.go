package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sjtucsn/wechatpp-backend/utils"
	"time"
)

type Chat struct {
	Id        uint
	Hash      string
	QuserId   string
	AuserId   string
	StartTime time.Time
	EndTime   time.Time
	Grade     int
	Token     int
	State     bool
}

func CreateChat(db *gorm.DB, QuserId string, AuserId string, token int) (chat Chat) {
	startTime := time.Now()
	s := QuserId + AuserId + startTime.String()
	hash := utils.CalHash(s)

	chat = Chat{
		Hash: hash,
		QuserId: QuserId,
		AuserId: AuserId,
		StartTime: startTime,
		Token: token,
	}

	// 创建
	db.Create(&chat)
	return
}

func EndChat(db *gorm.DB, hash string) (Chat) {
	var chat Chat
	db.Where("hash = ?", hash).First(&chat).Update("end_time", time.Now())
	return chat
}