package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sjtucsn/wechatpp-backend/utils"
	"os"
	"path"
	"time"
)

type Solution struct {
	Id uint
	Hash string
	ExamHash string
	Title string
	SolveId string
	CreateTime time.Time
	Like uint
	AccessIds string
}

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

	// 创建
	db.Create(&solution)
	imagePath := path.Join("public", "solutions", solution.Hash)
	if err := os.Mkdir(imagePath, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	return
}

func QuerySolutionsByExamHash(db *gorm.DB, examHash string) ([]Solution) {
	var solutions []Solution
	db.Where("exam_hash = ?", examHash).Find(&solutions)
	return solutions
}

func QuerySolutionsByHash(db *gorm.DB, hash string) (Solution) {
	var solution Solution
	db.Where("hash = ?", hash).First(&solution)
	return solution
}