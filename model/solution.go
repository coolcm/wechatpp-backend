package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sjtucsn/wechatpp-backend/utils"
	"os"
	"path"
	"time"
)

// 答题图片信息表
type Solution struct {
	Id         uint
	Hash       string    //答题过程详解图片的唯一认证，也作为寻址参数
	ExamHash   string    //所解答的考卷的hash值
	Title      string    //解答内容的摘要（图片名）
	SolveId    string    //试卷解答者id
	CreateTime time.Time //图片上传时间
	Like       uint      //该解答获得的赞同数
	AccessIds  string    //支付token有权限查看该答案的用户id
}

// 创建一条新的答题记录
func CreateSolution (db *gorm.DB, examHash string, solveId string, title string) (solution Solution) {
    createTime := time.Now()
    s := examHash + solveId + title + createTime.String()
	hash := utils.CalHash(s)

	solution = Solution{
		Hash:       hash,
		ExamHash:   examHash,
		Title:      title,
		SolveId:   solveId,
		CreateTime: createTime,
		AccessIds:  solveId,
	}

	db.Create(&solution)
	imagePath := path.Join("public", "solutions", solution.Hash)
	// 创建答题图片存储目录
	if err := os.Mkdir(imagePath, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	return
}

// 根据试卷哈希值查找所有相关解答记录
func QuerySolutionsByExamHash(db *gorm.DB, examHash string) ([]Solution) {
	var solutions []Solution
	db.Where("exam_hash = ?", examHash).Find(&solutions)
	return solutions
}

// 根据哈希值查找解答记录
func QuerySolutionsByHash(db *gorm.DB, hash string) (Solution) {
	var solution Solution
	db.Where("hash = ?", hash).First(&solution)
	return solution
}