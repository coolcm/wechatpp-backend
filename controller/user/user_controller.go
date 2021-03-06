package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sjtucsn/wechatpp-backend/model"
	"github.com/sjtucsn/wechatpp-backend/utils"
	"net/http"
	"strconv"
)

// 处理创建用户的请求（根据微信号）
func HandleCreateUser(c *gin.Context) {
	wechatId := c.PostForm("wechat_id")
	if !utils.VerifyParams(c, map[string]string{"wechat_id": wechatId}) {
		return
	}

	token, err := strconv.Atoi(c.PostForm("token"))
	if err != nil || token <= 0{
		fmt.Println("wrong token number")
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "msg": "wrong token number"})
		return
	}

	user := model.CreateUser(wechatId, token)
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": user})
}

// 处理用户下线请求
func HandleUserLogout(c *gin.Context) {
	wechatId := c.PostForm("wechat_id")
	if !utils.VerifyParams(c, map[string]string{"wechat_id": wechatId}) {
		return
	}

	user := model.LogoutUser(wechatId)

	if user.WechatId != "" {
		c.JSON(http.StatusOK, gin.H{"status": "success", "user": user})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "info": "user does not exist"})
	}
}

// 处理根据微信号查询用户的请求
func HandleQueryUser(c *gin.Context) {
	wechatId := c.Query("wechat_id")
	if !utils.VerifyParams(c, map[string]string{"wechat_id": wechatId}) {
		return
	}

	user := model.QueryUser(wechatId)

	if user.WechatId != "" {
		user.Qtime = user.Qtime / 10e8
		user.Atime = user.Atime / 10e8
		c.JSON(http.StatusOK, gin.H{"status": "success", "user": user})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "info": "user does not exist"})
	}
}

// 处理奖励用户token的请求
func HandleAwardUserToken(c *gin.Context) {
	wechatId := c.PostForm("wechat_id")
	if !utils.VerifyParams(c, map[string]string{"wechat_id": wechatId}) {
		return
	}

	token, err := strconv.Atoi(c.PostForm("token"))
	if err != nil || token <= 0{
		fmt.Println("wrong token number")
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "msg": "wrong token number"})
		return
	}

	user := model.QueryUser(wechatId)

	if user.WechatId != "" {
		user = model.AddUserToken(wechatId, token)
		c.JSON(http.StatusOK, gin.H{"status": "success", "user": user})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "info": "user does not exist"})
	}
}
