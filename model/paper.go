package model

import (
	"fmt"
	"github.com/sjtucsn/wechatpp-backend/utils"
	"os"
	"path"
	"time"
)

// 考卷记录表
type ExamPaper struct {
	Id          uint
	Hash        string    //试卷的唯一认证哈希值，同时也作为图片寻址参数
	Description string    //试卷描述
	Title       string    //文件名
	UploadId    string    // 试卷上传者的id
	UploadTime  time.Time //试卷上传时间
	PaperType   string    // 试卷类别
	Credit      uint      //试卷精华指数
}

// 增加一条考卷记录
func CreateExamPaper(description string, title string, uploadId string, paperType string) (paper ExamPaper) {
    uploadTime := time.Now()
    s := description + title + uploadId + paperType + uploadTime.String()
	hash := utils.CalHash(s)

	paper = ExamPaper{
		Hash: hash,
		Description: description,
		Title: title,
		UploadId: uploadId,
		UploadTime: uploadTime,
		PaperType: paperType,
	}

	Db.Create(&paper)
	imagePath := path.Join("public", "exams", paper.Hash)
	// 创建试卷存储目录
	if err := os.Mkdir(imagePath, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	return
}

// 根据试卷类别查询试卷
func QueryExamPaperByType(paperType string) ([]ExamPaper) {
	var paper []ExamPaper
	Db.Where("paper_type = ?", paperType).Find(&paper)
	return paper
}

// 根据试卷哈希值查询试卷
func QueryExamPaperByHash(hash string) (ExamPaper) {
	var paper ExamPaper
	Db.Where("hash = ?", hash).First(&paper)
	return paper
}

// 增加试卷的精华数
func AddExamPaperCredit(hash string) (ExamPaper){
	var paper ExamPaper
	Db.Where("hash = ?", hash).First(&paper)
	paper.Credit = paper.Credit + 1
	Db.Save(paper)
	return paper
}