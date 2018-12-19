package model

import (
	"fmt"
	"github.com/sjtucsn/wechatpp-backend/utils"
	"os"
	"path"
	"time"
)

// 答题图片信息表
type Solution struct {
	Id          uint
	Hash        string    //答题过程详解图片的唯一认证，也作为寻址参数
	ExamHash    string    //所解答的考卷的hash值
	Description string    //解答图片的描述
	Title       string    //解答图片的图片名
	SolveId     string    //试卷解答者id
	CreateTime  time.Time //图片上传时间
	Like        uint      //该解答获得的赞同数
	AccessIds   string    //支付token有权限查看该答案的用户id
}

// 创建一条新的答题记录
func CreateSolution (examHash string, solveId string, description string, title string) (solution Solution) {
    createTime := time.Now()
    s := examHash + solveId + description + title + createTime.String()
	hash := utils.CalHash(s)

	solution = Solution{
		Hash: hash,
		ExamHash: examHash,
		Description: description,
		Title: title,
		SolveId: solveId,
		CreateTime: createTime,
		AccessIds: solveId,
	}

	Db.Create(&solution)
	imagePath := path.Join("public", "solutions", solution.Hash)
	// 创建答题图片存储目录
	if err := os.Mkdir(imagePath, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	return
}

// 根据试卷哈希值查找所有相关解答记录
func QuerySolutionsByExamHash(examHash string) ([]Solution) {
	var solutions []Solution
	Db.Where("exam_hash = ?", examHash).Find(&solutions)
	return solutions
}

// 根据哈希值查找解答记录
func QuerySolutionsByHash(hash string) (Solution) {
	var solution Solution
	Db.Where("hash = ?", hash).First(&solution)
	return solution
}