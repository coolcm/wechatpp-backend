package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sjtucsn/wechatpp-backend/utils"
	"os"
	"path"
	"time"
)

type ExamPaper struct {
	Id uint
	Hash string
	Description string
	UploadId string
	UploadTime time.Time
	PaperType string
	credit uint
}

func CreateExamPaper(db *gorm.DB, description string, uploadId string, paperType string) (paper ExamPaper) {
    uploadTime := time.Now()
    s := description + uploadId + paperType + uploadTime.String()
	hash := utils.CalHash(s)

	paper = ExamPaper{
		Hash: hash,
		Description: description,
		UploadId: uploadId,
		UploadTime: uploadTime,
		PaperType: paperType,
	}

	// 创建
	db.Create(&paper)
	imagePath := path.Join("public", "exams", paper.Hash)
	if err := os.Mkdir(imagePath, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	return
}

func QueryExamPaper(db *gorm.DB, paperType string) ([]ExamPaper) {
	var paper []ExamPaper
	db.Where("paper_type = ?", paperType).Find(&paper)
	return paper
}

func QueryByHash(db *gorm.DB, hash string) (ExamPaper) {
	var paper ExamPaper
	db.Where("hash = ?", hash).First(&paper)
	return paper
}