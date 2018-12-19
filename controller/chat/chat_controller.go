package chat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sjtucsn/wechatpp-backend/model"
	"github.com/sjtucsn/wechatpp-backend/utils"
	"net/http"
	"strconv"
)

// 处理增加答疑记录请求
func HandleCreateChat(c *gin.Context) {
	QuserId := c.PostForm("from")
	AuserId := c.PostForm("to")
	if !utils.VerifyParams(c, map[string]string{"from": QuserId, "to": AuserId}) {
		return
	}
	if model.QueryUser(QuserId).Id == 0 || model.QueryUser(AuserId).Id == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "msg": "Quser or Auser does not exist"})
		return
	}

	chat := model.CreateChat(QuserId, AuserId)
	c.JSON(http.StatusOK, gin.H{"status": "success", "chat": chat})
}

// 处理查询答疑的请求
func HandleQueryChat(c *gin.Context) {
	hash := c.Query("hash")
	if !utils.VerifyParams(c, map[string]string{"hash": hash}) {
		return
	}

	chat := model.QueryChat(hash)
	if chat.Id != 0 {
		c.JSON(http.StatusOK, gin.H{"status": "success", "chat": chat})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "msg": "chat does not exist"})
	}
}

// 处理一条答疑已完成的请求
func HandleEndChat(c *gin.Context) {
	hash := c.PostForm("hash")
	if !utils.VerifyParams(c, map[string]string{"hash": hash}) {
		return
	}

	chat := model.QueryChat(hash)
	if chat.Id != 0 {
		chat.End()
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail"})
	}
}

// 处理给答疑打分的请求
func HandleScoreChat(c *gin.Context) {
	hash := c.PostForm("hash")
	if !utils.VerifyParams(c, map[string]string{"hash": hash}) {
		return
	}

	grade, err := strconv.Atoi(c.PostForm("grade"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "msg": "wrong grade number"})
		return
	}

	chat := model.QueryChat(hash)
	if chat.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "msg": "no such chat"})
	} else {
		chat.Score(grade)
		c.JSON(http.StatusOK, gin.H{"status": "success", "chat": chat})
	}
}
