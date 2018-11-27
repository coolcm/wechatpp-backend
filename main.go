package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sjtucsn/wechatpp-backend/controller/chat"
	"github.com/sjtucsn/wechatpp-backend/controller/paper"
	"github.com/sjtucsn/wechatpp-backend/controller/transaction"
	"github.com/sjtucsn/wechatpp-backend/controller/user"
	"github.com/sjtucsn/wechatpp-backend/model"
)

func setupRouter() *gin.Engine {

	r := gin.Default()

	// Get user value
	r.POST("/user/create", user.HandleCreateUser)
	r.GET("/user/query", user.HandleQueryUser)
	r.POST("/paper/upload", paper.HandleUploadExamPaper)
	r.GET("/paper/query", paper.HandleQueryExamPaper)
	r.GET("/paper/download", paper.HandleDownloadExamPaper)
	r.POST("/chat/create", chat.HandleCreateChat)
	r.GET("/chat/end", chat.HandleEndChat)
	r.POST("/transaction/create", transaction.HandleCreateTransaction)
	r.GET("/transaction/query", transaction.HandleQueryTransactionByHash)

	return r
}

func main() {
	defer func() {
		model.Db.Close()
		println("hahaha goodbye")
	}()
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}