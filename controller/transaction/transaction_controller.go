package transaction

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sjtucsn/wechatpp-backend/model"
	"net/http"
	"strconv"
)

// 处理创建交易的请求
func HandleCreateTransaction(c *gin.Context) {
	from := c.PostForm("from")
	to := c.PostForm("to")
	num, err := strconv.Atoi(c.PostForm("token"))
	Type := c.PostForm("type")
	if err != nil {
		fmt.Println("invalid token number")
	}
	transaction := model.CreateTransaction(model.Db, from, to, num, Type)
	c.JSON(http.StatusOK, gin.H{"status": "success", "transaction": transaction})
}

// 处理根据哈希标识查询交易的请求
func HandleQueryTransactionByHash(c *gin.Context)  {
	hash := c.Query("hash")
	transaction := model.QueryTransactionByHash(model.Db, hash)
	if transaction.Id != 0 {
		c.JSON(http.StatusOK, gin.H{"status": "success", "transaction": transaction})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": "transaction does not exist"})
	}
}