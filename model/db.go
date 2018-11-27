package model

import "github.com/jinzhu/gorm"

var Db *gorm.DB

func init() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&ExamPaper{})
	db.AutoMigrate(&Chat{})
	db.AutoMigrate(&Transaction{})
	Db = db
}