package paper

import (
	"github.com/gin-gonic/gin"
	"github.com/sjtucsn/wechatpp-backend/model"
	"net/http"
)

func HandleUploadExamPaper(c *gin.Context) {
	description := c.Query("description")
	uploadId := c.Query("wechat_id")
	paperType := c.Query("paperType")

	paper := model.CreateExamPaper(model.Db, description, uploadId, paperType)
	c.JSON(http.StatusOK, gin.H{"status": "success", "paper": paper})
}

func HandleQueryExamPaper(c *gin.Context) {
	paperType := c.Query("paperType")

	paper := model.QueryExamPaper(model.Db, paperType)

	if num := len(paper); num != 0 {
		c.JSON(http.StatusOK, gin.H{"status": "success", "size": num, "paper": paper})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "size": 0})
	}
}

