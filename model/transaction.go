package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sjtucsn/wechatpp-backend/utils"
	"time"
)

// token交易记录表
type Transaction struct {
	Id         uint
	Hash       string    //通过hash256计算得出的唯一哈希值作为交易的唯一认证
	To         string    //token去向用户的id
	From       string    //token来源用户的id
	Num        int       //交易的token数量
	CreateTime time.Time //交易创建时间
	Type       string    //支付类型:答疑(qa);查看解答(solution);奖励(award);充值(pay);消费券(spend)
}

// 创建一条新交易
func CreateTransaction(db *gorm.DB, from string, to string, num int, Type string) (Transaction) {
	createTime := time.Now()
	hash := utils.CalHash(from + to + string(num) + Type + createTime.String())
	transaction := Transaction{
		Hash: hash,
		From: from,
		To: to,
		Num: num,
		CreateTime: createTime,
		Type: Type,
	}
	db.Create(&transaction)
	return transaction
}

// 根据交易唯一哈希值查找交易
func QueryTransactionByHash(db *gorm.DB, hash string) (Transaction) {
	transaction := new(Transaction)
	db.Where("hash = ?", hash).First(transaction)
	return *transaction
}