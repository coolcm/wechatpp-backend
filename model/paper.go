package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sjtucsn/wechatpp-backend/utils"
	"os"
	"path"
	"time"
)

// 考卷记录表
type ExamPaper struct {
	Id          uint
	Hash        string    //试卷的唯一认证哈希值，同时也作为图片寻址参数
	Description string    //考卷内容描述（图片名）
	UploadId    string    // 试卷上传者的id
	UploadTime  time.Time //试卷上传时间
	PaperType   string    // 试卷类别
	credit      uint      //试卷精华指数
}

// 增加一条考卷记录
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

	db.Create(&paper)
	imagePath := path.Join("public", "exams", paper.Hash)
	// 创建试卷存储目录
	if err := os.Mkdir(imagePath, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	return
}

// 根据试卷类别查询试卷
func QueryExamPaper(db *gorm.DB, paperType string) ([]ExamPaper) {
	var paper []ExamPaper
	db.Where("paper_type = ?", paperType).Find(&paper)
	return paper
}

// 根据试卷哈希值查询试卷
func QueryByHash(db *gorm.DB, hash string) (ExamPaper) {
	var paper ExamPaper
	db.Where("hash = ?", hash).First(&paper)
	return paper
}