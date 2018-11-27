package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sjtucsn/wechatpp-backend/utils"
	"time"
)

type Transaction struct {
	Id uint
	Hash string
	To string
	From string
	Num int
	CreateTime time.Time
	Type string
}

func CreateTransaction(db *gorm.DB, from string, to string, num int, Type string) (Transaction) {
	createTime := time.Now()
	hash := utils.CalHash(from + to + string(num) + Type + createTime.String())
	transaction := Transaction{
		Hash:       hash,
		From:       from,
		To:         to,
		Num:        num,
		CreateTime: createTime,
		Type:       Type,
	}
	db.Create(&transaction)
	return transaction
}

func QueryTransactionByHash(db *gorm.DB, hash string) (Transaction) {
	transaction := new(Transaction)
	db.Where("hash = ?", hash).First(transaction)
	return *transaction
}