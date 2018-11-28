package model

import "github.com/jinzhu/gorm"

var Db *gorm.DB

func init() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &ExamPaper{}, &Chat{}, &Transaction{}, &Solution{})
	Db = db
}