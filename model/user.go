package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// 用户数据表
type User struct {
	Id           uint
	WechatId     string //即用户微信号作为用户身份id
	Token        int    //用户在系统中的积分
	PendingToken int    //尚未支付的token数
	Qtime        int    //用户作为提问者参与在线答疑的总时间
	Atime        int    //用户作为回答者参与在线答疑的总时间
	State        bool   //用户当前在线状态（是或否）
	Like         int    //用户通过解答考卷得到的赞同数总和
}

// 创建新用户
func CreateUser(db *gorm.DB, WechatId string, token int) (user User) {
	user = User{
		WechatId: WechatId,
		Token: token,
		State: true,
	}

	db.Create(&user)
	return
}

// 根据微信号查询用户
func QueryUser(db *gorm.DB, WechatId string) (User) {
	var user User
	db.Where("wechat_id = ?", WechatId).First(&user)
	return user
}