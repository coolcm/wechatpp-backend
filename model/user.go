package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	Id uint
	WechatId string
	Token int
	Qtime int
	Atime int
	State bool
	Like int
}

func CreateUser(db *gorm.DB, WechatId string, token int) (user User) {

	// Migrate the schema
	db.AutoMigrate(&User{})

	user = User{
		WechatId: WechatId,
		Token: token,
		State: true,
	}

	// 创建
	db.Create(&user)
	return
}

func QueryUser(db *gorm.DB, WechatId string) (User) {
	var user User
	db.Where("wechat_id = ?", WechatId).First(&user)
	return user
}