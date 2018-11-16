package paper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sjtucsn/wechatpp-backend/model"
	"io"
	"net/http"
	"os"
	"path"
)

func HandleUploadExamPaper(c *gin.Context) {
	uploadId := c.Query("wechat_id")
	paperType := c.Query("paper_type")
	pic, _ := c.FormFile("exam_paper")
	description := pic.Filename
	paper := model.CreateExamPaper(model.Db, description, uploadId, paperType)

	if err := c.SaveUploadedFile(pic, path.Join("public", "exams", paper.Hash, pic.Filename)); err != nil {
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "paper": paper})
	}
}

func HandleQueryExamPaper(c *gin.Context) {
	paperType := c.Query("paper_type")

	paper := model.QueryExamPaper(model.Db, paperType)

	if num := len(paper); num != 0 {
		c.JSON(http.StatusOK, gin.H{"status": "success", "size": num, "paper": paper})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "size": 0})
	}
}

func HandleDownloadExamPaper(c *gin.Context) {
	hash := c.Query("hash")

	paper := model.QueryByHash(model.Db, hash)
	if paper.Id == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "info": "exam paper does not exist"})
	} else {
		imagePath := path.Join("public", "exams", paper.Hash, paper.Description)
		if reader, err := os.Open(imagePath); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"status": "fail", "info": "exam paper does not exist"})
		} else {
			io.Copy(c.Writer, reader)
		}
	}
}
