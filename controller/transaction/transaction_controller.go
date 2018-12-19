package transaction

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sjtucsn/wechatpp-backend/model"
	"github.com/sjtucsn/wechatpp-backend/utils"
	"net/http"
	"strconv"
)

// 处理创建交易的请求
func HandleCreateTransaction(c *gin.Context) {
	from := c.PostForm("from")
	to := c.PostForm("to")
	num, err := strconv.Atoi(c.PostForm("token"))
	Type := c.PostForm("type")
	if err != nil || num == 0{
		fmt.Println("invalid token number")
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "msg": "token number can not be empty"})
		return
	}
	if !utils.VerifyParams(c, map[string]string{"from": from, "to": to, "type": Type}) {
		return
	}

	transaction := model.CreateTransaction(from, to, num, Type)
	c.JSON(http.StatusOK, gin.H{"status": "success", "transaction": transaction})
}

// 处理根据哈希标识查询交易的请求
func HandleQueryTransactionByHash(c *gin.Context)  {
	hash := c.Query("hash")
	if !utils.VerifyParams(c, map[string]string{"hash": hash}) {
		return
	}

	transaction := model.QueryTransactionByHash(hash)
	if transaction.Id != 0 {
		c.JSON(http.StatusOK, gin.H{"status": "success", "transaction": transaction})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": "transaction does not exist"})
	}
}