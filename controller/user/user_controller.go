package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sjtucsn/wechatpp-backend/model"
	"net/http"
	"strconv"
)

func HandleCreateUser(c *gin.Context) {
	wechatId := c.Query("wechat_id")
	token, err := strconv.Atoi(c.Query("token"))
	if err != nil {
		fmt.Println("wrong token number")
	}

	user := model.CreateUser(model.Db, wechatId, token)
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": user})
}

func HandleQueryUser(c *gin.Context) {
	wechatId := c.Query("wechat_id")

	user := model.QueryUser(model.Db, wechatId)

	if user.WechatId != "" {
		c.JSON(http.StatusOK, gin.H{"status": "success", "user": user})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "info": "user does not exist"})
	}
}
