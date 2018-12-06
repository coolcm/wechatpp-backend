package paper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sjtucsn/wechatpp-backend/model"
	"github.com/sjtucsn/wechatpp-backend/utils"
	"io"
	"net/http"
	"os"
	"path"
)

// 处理上传试卷请求
func HandleUploadExamPaper(c *gin.Context) {
	uploadId := c.PostForm("wechat_id")
	paperType := c.PostForm("paper_type")
	description := c.PostForm("description")
	if !utils.VerifyParams(c, map[string]string{"wechat_id": uploadId, "paper_type": paperType, "description": description}) {
		return
	}

	//获取post的试卷body
	if pic, err := c.FormFile("exam_paper"); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "msg": "couldn't get post exam paper"})
	} else {
		title := pic.Filename
		paper := model.CreateExamPaper(model.Db, description, title, uploadId, paperType)

		if err := c.SaveUploadedFile(pic, path.Join("public", "exams", paper.Hash, title)); err != nil {
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "success", "paper": paper})
		}
	}
}

// 处理根据试卷类型对试卷列表的请求
func HandleQueryExamByType(c *gin.Context) {
	paperType := c.Query("paper_type")
	if !utils.VerifyParams(c, map[string]string{"paper_type": paperType}) {
		return
	}

	paper := model.QueryExamPaperByType(model.Db, paperType)

	if num := len(paper); num != 0 {
		c.JSON(http.StatusOK, gin.H{"status": "success", "size": num, "paper": paper})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "size": 0})
	}
}

// 处理根据试卷哈希获取试卷内容和解答内容的请求
func HandleQueryExamAndSolutions(c *gin.Context) {
	paperHash := c.Query("hash")
	if !utils.VerifyParams(c, map[string]string{"hash": paperHash}) {
		return
	}

	paper := model.QueryExamPaperByHash(model.Db, paperHash)
	solutions := model.QuerySolutionsByExamHash(model.Db, paperHash)

	if paper.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "info": "exam paper does not exist"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "paper": paper, "solutions": solutions, "length": len(solutions)})
	}
}

// 处理根据哈希值下载试卷的请求
func HandleDownloadExamPaper(c *gin.Context) {
	hash := c.Query("hash")
	if !utils.VerifyParams(c, map[string]string{"hash": hash}) {
		return
	}

	paper := model.QueryExamPaperByHash(model.Db, hash)
	if paper.Id == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "info": "exam paper does not exist"})
	} else {
		imagePath := path.Join("public", "exams", paper.Hash, paper.Title)
		if reader, err := os.Open(imagePath); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"status": "fail", "info": "exam paper does not exist"})
		} else {
			io.Copy(c.Writer, reader)
		}
	}
}
